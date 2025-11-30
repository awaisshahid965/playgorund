// mod guess_number_game;
// mod power;
mod shape_geomatory_calculator;
use shape_geomatory_calculator::{
    geomatory::{ GeomatoryShape, Rectangle, Circle, Square },
    geomatory_traits::Geometry,
};

fn main() {
    // guess_number_game::guess_number_game();
    // println!("power fun return value is: {}", power::pow(2, -3));
    // let square_shape = GeomatoryShape::Square(2.0);
    // println!("Square shape has values: {}", square_shape);

    let rectangle_shape = GeomatoryShape::Rectangle(Rectangle {
        width: 20.0,
        height: 30.0,
    });
    println!("Rectangle shape has values: {:?}", rectangle_shape);
    println!("Rectangle area is: {}", rectangle_shape.area());
    println!("Rectangle perimeter is: {}", rectangle_shape.perimeter());

    let circle_shape = GeomatoryShape::Circle(Circle {
        radius: 20.0,
    });
    println!("Circle shape has values: {:?}", circle_shape);
    println!("Circle area is: {}", circle_shape.area());
    println!("Circle perimeter is: {}", circle_shape.perimeter());

    let squre_shape = GeomatoryShape::Square(Square {
        side_length: 20.0,
    });
    println!("Square shape has values: {:?}", squre_shape);
    println!("Square area is: {}", squre_shape.area());
    println!("Square perimeter is: {}", squre_shape.perimeter());
}
