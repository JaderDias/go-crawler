use v6;
use Data::Dump;
my @to_work = (1);
my %working = ();
my $max_parallel = 2;
while @to_work || %working {
    while @to_work && ((keys %working) < $max_parallel) {
        my $url = shift @to_work;
        my $promise = Promise.start({
            await Promise.anyof(
                my $timed_out = Promise.in($url / 100),
                my $request = Promise.start({

                    if $url %% 5 {
                        if $url %% 2 {
                            die '/0$/ fail';
                        }

                        await Promise.in($url / 10);
                    }

                    if $url < 1e4 {
                        [ $url * 2, ($url * 2) + 1 ];
                    }
                }),
            );
            CATCH {
                default {
                    say sprintf "%4d: $_", $url;
                }
            }

            if $timed_out {
                say sprintf '%4d: /5$/ time out', $url;
            }

            $request.result;
        });

        %working{$url} = $promise;
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
