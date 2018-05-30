package mpgaeroot

import (
	"github.com/strongo-games/matching-pennies/server-go/mpapp"
	"github.com/strongo/log"
	"github.com/strongo/bots-framework/hosts/appengine"
	"github.com/julienschmidt/httprouter"
)

func init() {
	log.AddLogger(gaehost.GaeLogger)

	httpRouter := httprouter.New()

	mpapp.InitApp(httpRouter, gaehost.GaeBotHost{})
}
