use v6;

my @to_work = (1);
my %working = ();
while @to_work || %working {
    while @to_work {
        my $id = shift @to_work;
        my $promise = Promise.start({
            if $id < 1e3 {
                [ $id * 2, ($id * 2) + 1 ];
            }
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
