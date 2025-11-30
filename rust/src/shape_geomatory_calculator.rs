/*
You need to implement a program that calculates the area and perimeter of various shapes, including rectangles and circles. The program should have the following features:

 - Define an enum Shape to represent different types of shapes (e.g., Rectangle, Circle).
 - Implement a trait Geometry with methods area() and perimeter() that must be implemented for each shape.
 - Define structs Rectangle and Circle to represent these shapes.
 - Implement methods for calculating the area and perimeter for each shape.
 - Use modules to organize your code: one module for shapes and another for traits.

*/

pub mod geomatory {
    #[derive(Debug)]
    pub enum GeomatoryShape {
        Circle(Circle),
        Rectangle(Rectangle),
        Square(Square),
    }

    #[derive(Debug)]
    pub struct Rectangle {
        pub width: f64,
        pub height: f64,
    }

    #[derive(Debug)]
    pub struct Circle {
        pub radius: f64,
    }

    #[derive(Debug)]
    pub struct Square {
        pub side_length: f64,
    }
}

pub mod geomatory_traits {
    use std::f64::consts;
    use super::geomatory;

    pub trait Geometry {
        fn perimeter(&self) -> f64;
        fn area(&self) -> f64;
    }

    impl Geometry for geomatory::Rectangle {
        fn area(&self) -> f64 {
            self.width * self.height
        }

        fn perimeter(&self) -> f64 {
            2.0 * (self.width + self.height)
        }
    }

    impl Geometry for geomatory::Circle {
        fn area(&self) -> f64 {
            consts::PI * self.radius * self.radius
        }

        fn perimeter(&self) -> f64 {
            2.0 * consts::PI * self.radius
        }
    }

    impl Geometry for geomatory::Square {
        fn area(&self) -> f64 {
            self.side_length  * self.side_length
        }

        fn perimeter(&self) -> f64 {
            4.0 * self.side_length
        }
    }

    impl Geometry for geomatory::GeomatoryShape {
        fn area(&self) -> f64 {
            match *self {
                geomatory::GeomatoryShape::Circle(ref circle) => circle.area(),
                geomatory::GeomatoryShape::Rectangle(ref rectangle) => rectangle.area(),
                geomatory::GeomatoryShape::Square(ref square) => square.area(),
            }
        }
    
        fn perimeter(&self) -> f64 {
            match *self {
                geomatory::GeomatoryShape::Circle(ref circle) => circle.perimeter(),
                geomatory::GeomatoryShape::Rectangle(ref rectangle) => rectangle.perimeter(),
                geomatory::GeomatoryShape::Square(ref square) => square.perimeter(),
            }
        }
    }
}

pub mod display_trait {
    use std::fmt;
    use super::geomatory::{ Circle, Rectangle };

    impl fmt::Display for Rectangle {
        fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
            write!(f, "Rectangle {{ width: {}, height: {} }}", self.width, self.height)
        }
    }

    impl fmt::Display for Circle {
        fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
            write!(f, "Circle {{ radius: {} }}", self.radius)
        }
    }
}

