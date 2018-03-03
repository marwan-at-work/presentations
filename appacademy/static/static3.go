package static

// START OMIT
type task struct {
	id int
}

type job string

func doIt(x task) job {
	return doItNow(x)
}

func doItNow(x task) job {
	// do some work and
	// return the job name.
	return ""
}

// END OMIT
