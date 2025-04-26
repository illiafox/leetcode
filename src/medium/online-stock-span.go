package medium

type StockSpanner struct {
	stack []stockSpannerElem
}

func StockSpannerConstructor() StockSpanner {
	return StockSpanner{}
}

type stockSpannerElem struct {
	price int
	span  int
}

func (s *StockSpanner) Next(price int) int {
	steps := 1
	for len(s.stack) > 0 && price >= s.stack[len(s.stack)-1].price {
		prev := s.stack[len(s.stack)-1]
		steps += prev.span

		s.stack = s.stack[:len(s.stack)-1]
	}

	s.stack = append(s.stack, stockSpannerElem{
		price: price,
		span:  steps,
	})

	return steps
}
