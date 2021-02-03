package main

import {
	types
}

func (mf *CustomerFact) GetWhatToSay(sentence string) string {
	return fmt.Sprintf("Let say \"%s\"", sentence)
}