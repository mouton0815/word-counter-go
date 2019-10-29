package main

type Gatherer interface {
    Gather()
}

type GathererImpl struct {
    pathQueue chan string
}

func (g GathererImpl) Gather() {
    // TODO: Dummy implementation, add a real one
    g.pathQueue <- "hello world"
    g.pathQueue <- "world hello hello"
    g.pathQueue <- "this world is my world"
    g.pathQueue <- "hello my world"
    close(g.pathQueue)
}

// Factory function
func NewGatherer(pathQueue chan string) Gatherer {
    return GathererImpl{ pathQueue: pathQueue }
}

