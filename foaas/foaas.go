package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"
)

var dataPlaceholder = ":<data>:"

func dict() map[string]string {
	result := map[string]string{
		"version":     "2.0.0",
		"anyway":      "Who the fuck are you anyway, :<data>:, why are you stirring up so much trouble, and, who pays you? - :<data>:",
		"asshole":     "Fuck you, asshole. - :<data>:",
		"awesome":     "This is Fucking Awesome. - :<data>:",
		"back":        ":<data>:, back the fuck off. - :<data>:",
		"bag":         "Eat a bag of fucking dicks. - :<data>:",
		"ballmer":     "Fucking :<data>: is a fucking pussy. I'm going to fucking bury that guy, I have done it before, and I will do it again. I'm going to fucking kill :<data>:. - :<data>:",
		"bday":        "Happy Fucking Birthday, :<data>:. - :<data>:",
		"because":     "Why? Because fuck you, that's why. - :<data>:",
		"blackadder":  ":<data>:, your head is as empty as a eunuch’s underpants. Fuck off! - :<data>:",
		"bm":          "Bravo mike, :name. - :<data>:",
		"bucket":      "Please choke on a bucket of cocks. - :<data>:",
		"bus":         "Christ on a bendy-bus, :<data>:, don't be such a fucking faff-arse. - :<data>:",
		"bye":         "Fuckity bye! - :<data>:",
		"caniuse":     "Can you use :tool? Fuck no! - :<data>:",
		"chainsaw":    "Fuck me gently with a chainsaw, :name. Do I look like Mother Teresa? - :<data>:",
		"cocksplat":   "Fuck off :<data>:, you worthless cocksplat - :<data>:",
		"cool":        "Cool story, bro. - :<data>:",
		"cup":         "How about a nice cup of shut the fuck up? - :<data>:",
		"dalton":      ":<data>:: A fucking problem solving super-hero. - :<data>:",
		"deraadt":     ":<data>: you are being the usual slimy hypocritical asshole... You may have had value ten years ago, but people will see that you don't anymore. - :<data>:",
		"diabetes":    "I'd love to stop and chat to you but I'd rather have type 2 diabetes. - :<data>:",
		"dosomething": ":<data>: the fucking :<data>:! - :<data>:",
		"donut":       ":<data>:, go and take a flying fuck at a rolling donut. - :<data>:",
		"equity":      "Equity only? long hours? zero pay? well :<data>:, just sign me right the fuck up. - :<data>:",
		"everyone":    "Everyone can go and fuck off. - :<data>:",
		"everything":  "Fuck everything. - :<data>:",
		"family":      "Fuck you, your whole family, your pets, and your feces. - :<data>:",
		"fascinating": "Fascinating story, in what chapter do you shut the fuck up? - :<data>:",
		"fewer":       "Go fuck yourself :<data>:, you'll disappoint fewer people. - :<data>:",
		"field":       "And :<data>: said unto :<data>:, 'verily, cast thine eyes upon the field in which i grow my fucks', and :<data>: gave witness unto the field, and saw that it was barren. - :<data>:",
		"flying":      "I don't give a flying fuck. - :<data>:",
		"ftfy":        "Fuck that, fuck you - :<data>:",
		"fts":         "Fuck that shit, :<data>:. - :<data>:",
		"fyyff":       "Fuck you, you fucking fuck. - :<data>:",
		"gfy":         "Golf foxtrot yankee, :<data>:. - :<data>:",
		"give":        "I give zero fucks. - :<data>:",
		"greed":       "The point is, ladies and gentleman, that :<data>: -- for lack of a better word -- is good. :<data>: is right. :<data>: works. :<data>: clarifies, cuts through, and captures the essence of the evolutionary spirit. :<data>:, in all of its forms -- :<data>: for life, for money, for love, knowledge -- has marked the upward surge of mankind. - :<data>:",
		"horse":       "Fuck you and the horse you rode in on. - :<data>:",
		"immensity":   "You can not imagine the immensity of the fuck i do not give. - :<data>:",
		"ing":         "Fucking fuck off, :<data>:. - :<data>:",
		"jinglebells": "Fuck you, fuck me, fuck your family. fuck your father, fuck your mother, fuck you and me. - :<data>:",
		"keep":        ":<data>:: fuck off. and when you get there, fuck off from there too. then fuck off some more. keep fucking off until you get back here. then fuck off again. - :<data>:",
		"keepcalm":    "Keep the fuck calm and :<data>:! - :<data>:",
		"king":        "Oh fuck off, just really fuck off you total dickface. christ, :<data>:, you are fucking thick. - :<data>:",
		"life":        "Fuck my life. - :<data>:",
		"linus":       ":<data>:, there aren't enough swear-words in the english language, so now i'll have to call you perkeleen vittupää just to express my disgust and frustration with this crap. - :<data>:",
		"logs":        "Check your fucking logs! - :<data>:",
		"look":        ":<data>:, do i look like i give a fuck? - :<data>:",
		"looking":     "Looking for a fuck to give. - :<data>:",
		"madison":     "What you've just said is one of the most insanely idiotic things i have ever heard, :<data>:. at no point in your rambling, incoherent response were you even close to anything that could be considered a rational thought. everyone in this room is now dumber for having listened to it. i award you no points :<data>:, and may god have mercy on your soul. - :<data>:",
		"maybe":       "Maybe. maybe not. maybe fuck yourself. - :<data>:",
		"me":          "Fuck me. - :<data>:",
		"mornin":      "Happy fuckin' mornin! - :<data>:",
		"no":          "No fucks given. - :<data>:",
		"nugget":      "Well :<data>:, aren't you a shining example of a rancid fuck-nugget. - :<data>:",
		"off":         "Fuck off, :<data>:. - :<data>:",
		"off-with":    "Fuck off with :<data>: - :<data>:",
		"outside":     ":<data>:, why don't you go outside and play hide-and-go-fuck-yourself? - :<data>:",
		"particular":  "Fuck this :<data>: in particular. - :<data>:",
		"pink":        "Well, fuck me pink. - :<data>:",
		"problem":     "What the fuck is your problem :<data>:? - :<data>:",
		"programmer":  "Fuck you, i'm a programmer, bitch! - :<data>:",
		"pulp":        ":<data>:, motherfucker, do you speak it? - :<data>:",
		"question":    "To fuck off, or to fuck off (that is not a question) - :<data>:",
		"ratsarse":    "I don't give a rat's arse. - :<data>:",
		"retard":      "You fucktard! - :<data>:",
		"ridiculous":  "That's fucking ridiculous - :<data>:",
		"rtfm":        "Read the fucking manual! - :<data>:",
		"sake":        "For fuck's sake! - :<data>:",
		"shakespeare": ":<data>:, thou clay-brained guts, thou knotty-pated fool, thou whoreson obscene greasy tallow-catch! - :<data>:",
		"shit":        "Fuck this shit! - :<data>:",
		"shutup":      ":<data>:, shut the fuck up. - :<data>:",
		"single":      "Not a single fuck was given. - :<data>:",
		"thanks":      "Fuck you very much. - :<data>:",
		"that":        "Fuck that. - :<data>:",
		"think":       ":<data>:, you think i give a fuck? - :<data>:",
		"thinking":    ":<data>:, what the fuck were you actually thinking? - :<data>:",
		"this":        "Fuck this. - :<data>:",
		"thumbs":      "Who has two thumbs and doesnt give a fuck? :<data>:. - :<data>:",
		"too":         "Thanks, fuck you too. - :<data>:",
		"tucker":      "Come the fuck in or fuck the fuck off. - :<data>:",
		"waste":       "I don't waste my fucking time with your bullshit :<data>:! - :<data>:",
		"what":        "What the fuck‽ - :<data>:",
		"xmas":        "Merry fucking christmas, :<data>:. - :<data>:",
		"yoda":        "Fuck off, you must, :<data>:. - :<data>:",
		"you":         "Fuck you, :<data>:. - :<data>:",
		"zayn":        "Ask me if i give a motherfuck ?!! - :<data>:",
		"zero":        "Zero, that's the number of fucks i give. - :<data>:",
	}
	return result
}

func getWord(key string) (string, error) {
	// fmt.Println("===> ", key)
	result, ok := dict()[key]
	if !ok {
		return "", errors.New("invalid key")
	}
	return result, nil
}

func formatWord(req *http.Request) string {
	args := strings.Split(req.URL.Path, "/")
	args = slices.DeleteFunc(args, func(x string) bool {
		return (len(x) <= 0)
	})
	result, err := getWord(args[0])
	if err != nil {
		return "Invalid route"
	}

	if len(args)-1 != strings.Count(result, dataPlaceholder) {
		return "Error wrong number of argument"
	} else {
		for _, arg := range args[1:] {
			result = strings.Replace(result, dataPlaceholder, arg, 1)
		}
		return result
	}
}

func FoaasAnswer(w http.ResponseWriter, req *http.Request) {
	ans := formatWord(req)
	fmt.Printf("%v", ans)
	fmt.Fprintln(w, ans)
}

func handleKeysAsFunc() {
	for key, words := range dict() {
		nargs := strings.Count(words, dataPlaceholder)
		var route strings.Builder
		route.WriteString("/")
		route.WriteString(key)
		for i := 0; i < nargs; i++ {
			route.WriteString("/*")
		}
		fmt.Println("Declaring ", route.String())
		http.HandleFunc(route.String(), FoaasAnswer)
	}

}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	port := flag.Int("port", 19050, "local http port")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	args := flag.Args()
	fmt.Println(args)
	fmt.Println("Running on port: ", *port)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	handleKeysAsFunc()

	// http.HandleFunc("/version", FoaasAnswer)
	// http.HandleFunc("/anyway/*/*", FoaasAnswer)
	// http.HandleFunc("/asshole/*", FoaasAnswer)
	// http.HandleFunc("/awesome/*", FoaasAnswer)
	// http.HandleFunc("/anyway/*/*", FoaasAnswer)
	// http.HandleFunc("/back/*/*", FoaasAnswer)
	// http.HandleFunc("/bag/*", FoaasAnswer)
	// http.HandleFunc("/ballmer/*/*/*", FoaasAnswer)
	// http.HandleFunc("/bday/*/*", FoaasAnswer)
	// http.HandleFunc("/because/*", FoaasAnswer)
	// http.HandleFunc("/blackadder/*/*", FoaasAnswer)
	// http.HandleFunc("/bm/*/*", FoaasAnswer)
	// http.HandleFunc("/bucket/*", FoaasAnswer)
	// http.HandleFunc("/bus/*/*", FoaasAnswer)

	portString := fmt.Sprintf(":%v", *port)
	http.ListenAndServe(portString, nil)

}
