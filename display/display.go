package display

import "island.com/mapcreater"

type Display interface{
	Show (*mapcreater.Island)
}