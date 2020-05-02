package main

type errConst string

func (e errConst) Error() string { return string(e) }

const errTooManyArgs = errConst("too many args")
