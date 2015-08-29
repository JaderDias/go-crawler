use v6;

my $supply = Supply.new;
$supply.tap: -> $id {
    say "$id";
    if $id < 1e3 {
        $supply.emit($id * 2);
        $supply.emit(($id * 2) + 1);
    }
}

$supply.emit(1);
