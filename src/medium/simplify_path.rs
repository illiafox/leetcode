struct Solution;

impl Solution {
    pub fn simplify_path(path: String) -> String {
        let mut stack = vec![];

        for token in path.split('/').filter(|s| !s.is_empty()) {
            match token {
                "." | "/" => {}
                ".." => _ = stack.pop(),
                s => stack.push(s),
            }
        }

        format!("/{}", stack.join("/"))
    }
}

#[test]
fn test() {
    let test_cases = [
        ("/a/./b/../../c/", "/c"),
        ("/home/", "/home"),
        ("/home//foo/", "/home/foo"),
        ("/home/user/Documents/../Pictures", "/home/user/Pictures"),
        ("/../", "/"),
        ("/.../a/../b/c/../d/./", "/.../b/d"),
    ];

    for (path, expected) in test_cases {
        assert_eq!(
            Solution::simplify_path(path.to_string()),
            expected,
            "failed for input: nums = {:?}",
            path
        );
    }
}
