// Rust 的结构体没有继承, trait类似于接口，也可以是

#[cfg(test)]
mod test {

    struct Teacher {
        name: String,
    }

    struct Student {
        name: String,
    }

    trait SayAble {
        fn say(&self) -> String;
    }

    impl SayAble for Teacher {
        fn say(&self) -> String {
            let mut say_str = String::from("Teacher Say ");
            say_str.push_str(&self.name);
            say_str
        }
    }
    impl SayAble for Student {
        fn say(&self) -> String {
            let mut say_str = String::from("Student Say ");
            say_str.push_str(&self.name);
            say_str
        }
    }

    #[test]
    fn test_struct() {
        let t = Teacher {
            name: "Mr.Chen".to_string(),
        };
        assert_eq!(t.say(), "Teacher Say Mr.Chen");
        let s = Student {
            name: "Chess".to_string(),
        };
        s.say();
        assert_eq!(s.say(), "Student Say Chess");
    }
}
