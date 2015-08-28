use v6;
my @to_work = ('url');
my @working = ();
my $max_network_connections = 4;
my $timeout_s = 1;
while @to_work || @working {
    my $url = shift @to_work;
    if $url {
        my $promise = Promise.start({
            await Promise.anyof(
                my $timed_out = Promise.in($timeout_s),
                my $request = Promise.start({
                    # there's 10% of chance of a request failing
                    die if $url ne 'url' && rand < .1;
                    CATCH {
                        default {
                            say "$url failed";
                        }
                    }

                    # there's 10% chance of a request timing out;
                    await Promise.in($timeout_s * 2) if $url ne 'url' && rand < .1;

                    [ $url ~ 0, $url ~ 1 ];
                }),
            );
        
            if $timed_out {
                say "$url timed out";
                [];
            } else { 
                $request.result;
            }
        });

        push @working, $promise;
    }

    if !@to_work || @working >= $max_network_connections {
        await Promise.anyof(@working);
        for @working -> $promise {
            if ($promise) {
                push @to_work, @($promise.result);
            }
        }

        @working .= grep({ !$_ });
    }
}
