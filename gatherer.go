package main

type Gatherer interface {
    Gather(pathQueue chan string)
}

type GathererImpl struct {

}

func (g *GathererImpl) Gather(pathQueue chan string) {
    pathQueue <- "hello world"
    pathQueue <- "world hello hello"
    pathQueue <- "this world is my world"
    pathQueue <- "hello my world"
    close(pathQueue)
}

// Factory function
func NewGatherer() Gatherer {
    return &GathererImpl{}
}

