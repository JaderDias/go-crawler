use v6;

my @supplies = (Supply.new, Supply.new);
for @supplies -> $supply {
    $supply.act: -> $id {
        say "$id";
        if $id < 1e3 {
            for (0, 1) -> $i {
                @supplies[$i].emit(($id * 2) + $i);
            }
        }
    }
}

@supplies[0].emit(1);
