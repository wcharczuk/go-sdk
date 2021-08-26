/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package slant

// Slant is the default font.
var Slant = Font{
	Height:		6,	// the nominal letter height in rows.
	Baseline:	5,	// the baseline does something cool with font alightment but ???.
	Width:		16,	// this is mostly unused, but kept for legacy reasons.
	Hardblank:	'$',
	Letters: [][]string{
		{
			"     $$",
			"    $$ ",
			"   $$  ",
			"  $$   ",
			" $$    ",
			"$$     ",
		},
		{
			"    __",
			"   / /",
			"  / / ",
			" /_/  ",
			"(_)   ",
			"      ",
		},
		{
			" _ _ ",
			"( | )",
			"|/|/ ",
			" $   ",
			"$    ",
			"     ",
		},
		{
			"     __ __ ",
			"  __/ // /_",
			" /_  _  __/",
			"/_  _  __/ ",
			" /_//_/    ",
			"           ",
		},
		{
			"     __",
			"   _/ /",
			"  / __/",
			" (_  ) ",
			"/  _/  ",
			"/_/    ",
		},
		{
			"   _   __",
			"  (_)_/_/",
			"   _/_/  ",
			" _/_/_   ",
			"/_/ (_)  ",
			"         ",
		},
		{
			"   ___   ",
			"  ( _ )  ",
			" / __ \\/|",
			"/ /_/  < ",
			"\\____/\\/ ",
			"         ",
		},
		{
			"  _ ",
			" ( )",
			" |/ ",
			" $  ",
			"$   ",
			"    ",
		},
		{
			"     __",
			"   _/_/",
			"  / /  ",
			" / /   ",
			"/ /    ",
			"|_|    ",
		},
		{
			"     _ ",
			"    | |",
			"    / /",
			"   / / ",
			" _/_/  ",
			"/_/    ",
		},
		{
			"       ",
			"  __/|_",
			" |    /",
			"/_ __| ",
			" |/    ",
			"       ",
		},
		{
			"       ",
			"    __ ",
			" __/ /_",
			"/_  __/",
			" /_/   ",
			"       ",
		},
		{
			"   ",
			"   ",
			"   ",
			" _ ",
			"( )",
			"|/ ",
		},
		{
			"       ",
			"       ",
			" ______",
			"/_____/",
			"  $    ",
			"       ",
		},
		{
			"   ",
			"   ",
			"   ",
			" _ ",
			"(_)",
			"   ",
		},
		{
			"       __",
			"     _/_/",
			"   _/_/  ",
			" _/_/    ",
			"/_/      ",
			"         ",
		},
		{
			"   ____ ",
			"  / __ \\",
			" / / / /",
			"/ /_/ / ",
			"\\____/  ",
			"        ",
		},
		{
			"   ___",
			"  <  /",
			"  / / ",
			" / /  ",
			"/_/   ",
			"      ",
		},
		{
			"   ___ ",
			"  |__ \\",
			"  __/ /",
			" / __/ ",
			"/____/ ",
			"       ",
		},
		{
			"   _____",
			"  |__  /",
			"   /_ < ",
			" ___/ / ",
			"/____/  ",
			"        ",
		},
		{
			"   __ __",
			"  / // /",
			" / // /_",
			"/__  __/",
			"  /_/   ",
			"        ",
		},
		{
			"    ______",
			"   / ____/",
			"  /___ \\  ",
			" ____/ /  ",
			"/_____/   ",
			"          ",
		},
		{
			"   _____",
			"  / ___/",
			" / __ \\ ",
			"/ /_/ / ",
			"\\____/  ",
			"        ",
		},
		{
			" _____",
			"/__  /",
			"  / / ",
			" / /  ",
			"/_/   ",
			"      ",
		},
		{
			"   ____ ",
			"  ( __ )",
			" / __  |",
			"/ /_/ / ",
			"\\____/  ",
			"        ",
		},
		{
			"   ____ ",
			"  / __ \\",
			" / /_/ /",
			" \\__, / ",
			"/____/  ",
			"        ",
		},
		{
			"     ",
			"   _ ",
			"  (_)",
			" _   ",
			"(_)  ",
			"     ",
		},
		{
			"     ",
			"   _ ",
			"  (_)",
			" _   ",
			"( )  ",
			"|/   ",
		},
		{
			"  __",
			" / /",
			"/ / ",
			"\\ \\ ",
			" \\_\\",
			"    ",
		},
		{
			"       ",
			"  _____",
			" /____/",
			"/____/ ",
			"  $    ",
			"       ",
		},
		{
			"__  ",
			"\\ \\ ",
			" \\ \\",
			" / /",
			"/_/ ",
			"    ",
		},
		{
			"  ___ ",
			" /__ \\",
			"  / _/",
			" /_/  ",
			"(_)   ",
			"      ",
		},
		{
			"   ______ ",
			"  / ____ \\",
			" / / __ `/",
			"/ / /_/ / ",
			"\\ \\__,_/  ",
			" \\____/   ",
		},
		{
			"    ___ ",
			"   /   |",
			"  / /| |",
			" / ___ |",
			"/_/  |_|",
			"        ",
		},
		{
			"    ____ ",
			"   / __ )",
			"  / __  |",
			" / /_/ / ",
			"/_____/  ",
			"         ",
		},
		{
			"   ______",
			"  / ____/",
			" / /     ",
			"/ /___   ",
			"\\____/   ",
			"         ",
		},
		{
			"    ____ ",
			"   / __ \\",
			"  / / / /",
			" / /_/ / ",
			"/_____/  ",
			"         ",
		},
		{
			"    ______",
			"   / ____/",
			"  / __/   ",
			" / /___   ",
			"/_____/   ",
			"          ",
		},
		{
			"    ______",
			"   / ____/",
			"  / /_    ",
			" / __/    ",
			"/_/       ",
			"          ",
		},
		{
			"   ______",
			"  / ____/",
			" / / __  ",
			"/ /_/ /  ",
			"\\____/   ",
			"         ",
		},
		{
			"    __  __",
			"   / / / /",
			"  / /_/ / ",
			" / __  /  ",
			"/_/ /_/   ",
			"          ",
		},
		{
			"    ____",
			"   /  _/",
			"   / /  ",
			" _/ /   ",
			"/___/   ",
			"        ",
		},
		{
			"       __",
			"      / /",
			" __  / / ",
			"/ /_/ /  ",
			"\\____/   ",
			"         ",
		},
		{
			"    __ __",
			"   / //_/",
			"  / ,<   ",
			" / /| |  ",
			"/_/ |_|  ",
			"         ",
		},
		{
			"    __ ",
			"   / / ",
			"  / /  ",
			" / /___",
			"/_____/",
			"       ",
		},
		{
			"    __  ___",
			"   /  |/  /",
			"  / /|_/ / ",
			" / /  / /  ",
			"/_/  /_/   ",
			"           ",
		},
		{
			"    _   __",
			"   / | / /",
			"  /  |/ / ",
			" / /|  /  ",
			"/_/ |_/   ",
			"          ",
		},
		{
			"   ____ ",
			"  / __ \\",
			" / / / /",
			"/ /_/ / ",
			"\\____/  ",
			"        ",
		},
		{
			"    ____ ",
			"   / __ \\",
			"  / /_/ /",
			" / ____/ ",
			"/_/      ",
			"         ",
		},
		{
			"   ____ ",
			"  / __ \\",
			" / / / /",
			"/ /_/ / ",
			"\\___\\_\\ ",
			"        ",
		},
		{
			"    ____ ",
			"   / __ \\",
			"  / /_/ /",
			" / _, _/ ",
			"/_/ |_|  ",
			"         ",
		},
		{
			"   _____",
			"  / ___/",
			"  \\__ \\ ",
			" ___/ / ",
			"/____/  ",
			"        ",
		},
		{
			"  ______",
			" /_  __/",
			"  / /   ",
			" / /    ",
			"/_/     ",
			"        ",
		},
		{
			"   __  __",
			"  / / / /",
			" / / / / ",
			"/ /_/ /  ",
			"\\____/   ",
			"         ",
		},
		{
			" _    __",
			"| |  / /",
			"| | / / ",
			"| |/ /  ",
			"|___/   ",
			"        ",
		},
		{
			" _       __",
			"| |     / /",
			"| | /| / / ",
			"| |/ |/ /  ",
			"|__/|__/   ",
			"           ",
		},
		{
			"   _  __",
			"  | |/ /",
			"  |   / ",
			" /   |  ",
			"/_/|_|  ",
			"        ",
		},
		{
			"__  __",
			"\\ \\/ /",
			" \\  / ",
			" / /  ",
			"/_/   ",
			"      ",
		},
		{
			" _____",
			"/__  /",
			"  / / ",
			" / /__",
			"/____/",
			"      ",
		},
		{
			"     ___",
			"    / _/",
			"   / /  ",
			"  / /   ",
			" / /    ",
			"/__/    ",
		},
		{
			"__    ",
			"\\ \\   ",
			" \\ \\  ",
			"  \\ \\ ",
			"   \\_\\",
			"      ",
		},
		{
			"     ___",
			"    /  /",
			"    / / ",
			"   / /  ",
			" _/ /   ",
			"/__/    ",
		},
		{
			"  //|",
			" |/||",
			"  $  ",
			" $   ",
			"$    ",
			"     ",
		},
		{
			"       ",
			"       ",
			"       ",
			"       ",
			" ______",
			"/_____/",
		},
		{
			"  _ ",
			" ( )",
			"  V ",
			" $  ",
			"$   ",
			"    ",
		},
		{
			"        ",
			"  ____ _",
			" / __ `/",
			"/ /_/ / ",
			"\\__,_/  ",
			"        ",
		},
		{
			"    __  ",
			"   / /_ ",
			"  / __ \\",
			" / /_/ /",
			"/_.___/ ",
			"        ",
		},
		{
			"       ",
			"  _____",
			" / ___/",
			"/ /__  ",
			"\\___/  ",
			"       ",
		},
		{
			"       __",
			"  ____/ /",
			" / __  / ",
			"/ /_/ /  ",
			"\\__,_/   ",
			"         ",
		},
		{
			"      ",
			"  ___ ",
			" / _ \\",
			"/  __/",
			"\\___/ ",
			"      ",
		},
		{
			"    ____",
			"   / __/",
			"  / /_  ",
			" / __/  ",
			"/_/     ",
			"        ",
		},
		{
			"         ",
			"   ____ _",
			"  / __ `/",
			" / /_/ / ",
			" \\__, /  ",
			"/____/   ",
		},
		{
			"    __  ",
			"   / /_ ",
			"  / __ \\",
			" / / / /",
			"/_/ /_/ ",
			"        ",
		},
		{
			"    _ ",
			"   (_)",
			"  / / ",
			" / /  ",
			"/_/   ",
			"      ",
		},
		{
			"       _ ",
			"      (_)",
			"     / / ",
			"    / /  ",
			" __/ /   ",
			"/___/    ",
		},
		{
			"    __  ",
			"   / /__",
			"  / //_/",
			" / ,<   ",
			"/_/|_|  ",
			"        ",
		},
		{
			"    __",
			"   / /",
			"  / / ",
			" / /  ",
			"/_/   ",
			"      ",
		},
		{
			"            ",
			"   ____ ___ ",
			"  / __ `__ \\",
			" / / / / / /",
			"/_/ /_/ /_/ ",
			"            ",
		},
		{
			"        ",
			"   ____ ",
			"  / __ \\",
			" / / / /",
			"/_/ /_/ ",
			"        ",
		},
		{
			"       ",
			"  ____ ",
			" / __ \\",
			"/ /_/ /",
			"\\____/ ",
			"       ",
		},
		{
			"         ",
			"    ____ ",
			"   / __ \\",
			"  / /_/ /",
			" / .___/ ",
			"/_/      ",
		},
		{
			"        ",
			"  ____ _",
			" / __ `/",
			"/ /_/ / ",
			"\\__, /  ",
			"  /_/   ",
		},
		{
			"        ",
			"   _____",
			"  / ___/",
			" / /    ",
			"/_/     ",
			"        ",
		},
		{
			"        ",
			"   _____",
			"  / ___/",
			" (__  ) ",
			"/____/  ",
			"        ",
		},
		{
			"   __ ",
			"  / /_",
			" / __/",
			"/ /_  ",
			"\\__/  ",
			"      ",
		},
		{
			"        ",
			"  __  __",
			" / / / /",
			"/ /_/ / ",
			"\\__,_/  ",
			"        ",
		},
		{
			"       ",
			" _   __",
			"| | / /",
			"| |/ / ",
			"|___/  ",
			"       ",
		},
		{
			"          ",
			" _      __",
			"| | /| / /",
			"| |/ |/ / ",
			"|__/|__/  ",
			"          ",
		},
		{
			"        ",
			"   _  __",
			"  | |/_/",
			" _>  <  ",
			"/_/|_|  ",
			"        ",
		},
		{
			"         ",
			"   __  __",
			"  / / / /",
			" / /_/ / ",
			" \\__, /  ",
			"/____/   ",
		},
		{
			"     ",
			" ____",
			"/_  /",
			" / /_",
			"/___/",
			"     ",
		},
		{
			"     __",
			"   _/_/",
			" _/_/  ",
			"< <    ",
			"/ /    ",
			"\\_\\    ",
		},
		{
			"     __",
			"    / /",
			"   / / ",
			"  / /  ",
			" / /   ",
			"/_/    ",
		},
		{
			"     _ ",
			"    | |",
			"    / /",
			"   _>_>",
			" _/_/  ",
			"/_/    ",
		},
		{
			"  /\\//",
			" //\\/ ",
			"  $   ",
			" $    ",
			"$     ",
			"      ",
		},
	},
}
