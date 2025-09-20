use std::cmp::Ordering;
use std::collections::{BinaryHeap, HashMap};

// https://leetcode.com/problems/design-a-food-rating-system

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct Food {
    pub rating: i32,
    pub name: String,
}

impl Ord for Food {
    fn cmp(&self, other: &Self) -> Ordering {
        // if there is a tie, return the item with the lexicographically smaller name.
        self.rating
            .cmp(&other.rating)
            .then_with(|| other.name.cmp(&self.name))
    }
}
impl PartialOrd for Food {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

struct FoodRatings {
    food_to_cuisine: HashMap<String, (String, i32)>,
    cuisine_rating: HashMap<String, BinaryHeap<Food>>,
}

impl FoodRatings {
    fn new(foods: Vec<String>, cuisines: Vec<String>, ratings: Vec<i32>) -> Self {
        let mut food_to_cuisine = HashMap::with_capacity(foods.len());
        let mut cuisine_rating: HashMap<String, BinaryHeap<Food>> = HashMap::new();

        for ((food, cuisine), rating) in foods.into_iter().zip(cuisines).zip(ratings) {
            food_to_cuisine.insert(food.clone(), (cuisine.clone(), rating));

            cuisine_rating
                .entry(cuisine)
                .or_default()
                .push(Food { name: food, rating });
        }

        FoodRatings {
            food_to_cuisine,
            cuisine_rating,
        }
    }

    fn change_rating(&mut self, food: String, new_rating: i32) {
        let cuisine = {
            let (c, cur) = self.food_to_cuisine.get_mut(&food).unwrap();
            *cur = new_rating;
            c.clone()
        };

        self.cuisine_rating.entry(cuisine).or_default().push(Food {
            name: food,
            rating: new_rating,
        });
    }

    fn highest_rated(&mut self, cuisine: String) -> String {
        let heap = match self.cuisine_rating.get_mut(&cuisine) {
            Some(h) => h,
            None => return String::new(),
        };

        while let Some(top) = heap.peek() {
            if let Some((_, cur_rating)) = self.food_to_cuisine.get(&top.name) {
                if *cur_rating == top.rating {
                    return top.name.clone();
                }
            }
            heap.pop();
        }
        String::new()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn build() -> FoodRatings {
        // foods:   name      cuisine   rating
        let foods = vec!["sushi", "ramen", "udon", "burger", "burrito"]
            .into_iter()
            .map(|s| s.to_string())
            .collect::<Vec<_>>();
        let cuisines = vec!["jp", "jp", "jp", "us", "mx"]
            .into_iter()
            .map(|s| s.to_string())
            .collect::<Vec<_>>();
        let ratings = vec![10, 8, 9, 7, 9];
        FoodRatings::new(foods, cuisines, ratings)
    }

    #[test]
    fn builds_and_queries() {
        let mut fr = build();
        assert_eq!(fr.highest_rated("jp".to_string()), "sushi");
        assert_eq!(fr.highest_rated("us".to_string()), "burger");
        assert_eq!(fr.highest_rated("mx".to_string()), "burrito");
    }

    #[test]
    fn tie_breaks_by_lex_name_ascending() {
        let foods = vec!["banana", "apple"]
            .into_iter()
            .map(|s| s.to_string())
            .collect();
        let cuisines = vec!["c", "c"].into_iter().map(|s| s.to_string()).collect();
        let ratings = vec![5, 5];
        let mut fr = FoodRatings::new(foods, cuisines, ratings);
        // at equal rating, lexicographically smaller name ("apple") should win
        assert_eq!(fr.highest_rated("c".to_string()), "apple");
    }

    #[test]
    fn change_rating_updates_and_lazy_deletes() {
        let mut fr = build();
        // initially jp top is sushi(10)
        assert_eq!(fr.highest_rated("jp".to_string()), "sushi");

        // boost ramen above sushi
        fr.change_rating("ramen".to_string(), 12);
        assert_eq!(fr.highest_rated("jp".to_string()), "ramen");

        // drop ramen below udon
        fr.change_rating("ramen".to_string(), 8);
        assert_eq!(fr.highest_rated("jp".to_string()), "sushi"); // still 10 vs udon 9 vs ramen 8

        // raise udon to tie sushi (10) — name tie-break => "sushi" < "udon", so sushi remains
        fr.change_rating("udon".to_string(), 10);
        assert_eq!(fr.highest_rated("jp".to_string()), "sushi");

        // raise udon to 11, now udon wins
        fr.change_rating("udon".to_string(), 11);
        assert_eq!(fr.highest_rated("jp".to_string()), "udon");
    }

    #[test]
    fn multiple_cuisines_are_isolated() {
        let mut fr = build();
        fr.change_rating("burger".to_string(), 100);
        assert_eq!(fr.highest_rated("us".to_string()), "burger");
        // jp unaffected
        assert_eq!(fr.highest_rated("jp".to_string()), "sushi");
    }
}
