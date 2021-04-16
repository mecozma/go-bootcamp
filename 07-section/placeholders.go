package main

// for keeping things easy to read and type-safe
type placeholder [5]string

var (
	zero = placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one = placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two = placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three = placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four = placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five = placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six = placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven = placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight = placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine = placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	separator = placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	dot = placeholder{
		"   ",
		"   ",
		"   ",
		"   ",
		" ░ ",
	}

	alarm = [...]placeholder{
		{
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
		{
			"███",
			"█ █",
			"███",
			"█ █",
			"█ █",
		},
		{
			"█  ",
			"█  ",
			"█  ",
			"█  ",
			"███",
		},
		{
			"███",
			"█ █",
			"███",
			"█ █",
			"█ █",
		},
		{
			"██ ",
			"█ █",
			"██ ",
			"█ █",
			"█ █",
		},
		{
			"█ █",
			"███",
			"█ █",
			"█ █",
			"█ █",
		},
		{
			" █ ",
			" █ ",
			" █ ",
			"   ",
			" █ ",
		},
		{
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
	}

	// array to hold all the digits.
	digits = [...]placeholder{zero, one, two, three, four, five, six, seven, eight, nine}
)
