use v6;

my @to_work = (1);
my %working = ();
my $max_parallel = 2;
while @to_work || %working {
    while @to_work && ((keys %working) < $max_parallel) {
        my $id = shift @to_work;
        my $promise = Promise.start({
            await Promise.anyof(
                my $timed_out = Promise.in($id / 100),
                my $request = Promise.start({

                    if $id %% 5 {
                        if $id %% 2 {
                            die '/0$/ fail';
                        }

                        await Promise.in($id / 50);
                    }

                    if $id < 1e3 {
                        [ $id * 2, ($id * 2) + 1 ];
                    }
                }),
            );
            CATCH {
                default {
                    say sprintf "%4d: $_", $id;
                }
            }

            if $timed_out {
                say sprintf '%4d: /5$/ time out', $id;
            }

            $request.result;
        });

        %working{$id} = $promise;
    }

    say "\t\t\twaiting " ~ join ', ', map { sprintf("%4d", $_) }, sort keys %working;
    await Promise.anyof(values %working);
    for values %working -> $promise {
        if ($promise && $promise.result) {
            push @to_work, @($promise.result);
        }
    }

    %working = map { $_ => %working{$_} }, grep { !%working{$_} }, keys %working;
}
