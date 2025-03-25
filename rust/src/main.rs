mod algorithm;
mod data_structure;
mod learn_rust;
use algorithm::main as algorithm_main;
use data_structure::main as data_structure_main;
use learn_rust::main as learn_rust_main;
fn main() {
    println!("Hello, Rust Algorithm ~");
    learn_rust_main();
    data_structure_main();
    algorithm_main();
}
