// https://hub.packtpub.com/concurrency-and-parallelism-in-golang-tutorial/

package main

import (
	"sync"
	"testing"
)

func Test_makeHotelReservation(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makeHotelReservation(tt.args.wg)
		})
	}
}

func Test_bookFlightTickets(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bookFlightTickets(tt.args.wg)
		})
	}
}

func Test_orderADress(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderADress(tt.args.wg)
		})
	}
}

func Test_payCreditCardBills(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payCreditCardBills(tt.args.wg)
		})
	}
}

func Test_writeAMail(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeAMail(tt.args.wg)
		})
	}
}

func Test_continueWritingMail1(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			continueWritingMail1(tt.args.wg)
		})
	}
}

func Test_continueWritingMail2(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			continueWritingMail2(tt.args.wg)
		})
	}
}

func Test_listenToAudioBook(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listenToAudioBook(tt.args.wg)
		})
	}
}

func Test_continueListeningToAudioBook(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			continueListeningToAudioBook(tt.args.wg)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
