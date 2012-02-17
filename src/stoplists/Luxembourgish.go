package gojustext

var LuxembourgishStoplist = map[string]bool {"an": true,
"der": true,
"de": true,
"vun": true,
"ass": true,
"a": true,
"am": true,
"den": true,
"vum": true,
"eng": true,
"déi": true,
"huet": true,
"zu": true,
"op": true,
"vu": true,
"De": true,
"e": true,
"mat": true,
"fir": true,
"war": true,
"dem": true,
"sech": true,
"gouf": true,
"och": true,
"en": true,
"Den": true,
"Joerhonnert": true,
"ze": true,
"Joer": true,
"et": true,
"aus": true,
"enger": true,
"als": true,
"engem": true,
"bis": true,
"hien": true,
"Canton": true,
"si": true,
"franséisch": true,
"bei": true,
"oder": true,
"wéi": true,
"gëtt": true,
"deen": true,
"net": true,
"mam": true,
"méi": true,
"un": true,
"um": true,
"dat": true,
"sinn": true,
"duerch": true,
"Departement": true,
"datt": true,
"ginn.": true,
"Gemeng": true,
"no": true,
"gebuer": true,
"Am": true,
"se": true,
"Arrondissement": true,
"nach": true,
"Regioun": true,
"hie": true,
"seng": true,
"v.": true,
"Déi": true,
"awer": true,
"Gemengen": true,
"administrativ": true,
"Andeelung": true,
"Si": true,
"zum": true,
"huet.": true,
"iwwer": true,
"Lëtzebuerg": true,
"ginn": true,
"Kanton.": true,
"hunn": true,
"Et": true,
"een": true,
"Säit": true,
"dee": true,
"An": true,
"wou": true,
"Numm": true,
"Chr.": true,
"Hien": true,
"do": true,
"ee": true,
"Stad": true,
"gestuerwen": true,
"hir": true,
"deem": true,
"hat": true,
"befaasst": true,
"sengem": true,
"\"Dës": true,
"ass.": true,
"läit": true,
"__NOTOC__": true,
"huet,": true,
"des": true,
"Hie": true,
"senger": true,
"där": true,
"all": true,
"ginn,": true,
"ass,": true,
"la": true,
"Uertschaft": true,
"wat": true,
"Zäit": true,
"tëscht": true,
"Mee": true,
"säi": true,
"Deel": true,
"goufen": true,
"ëm": true,
"Um": true,
"du": true,
"gouf.": true,
"ënner": true,
"E": true,
"ganz": true,
"well": true,
"hu": true,
"virun": true,
"Vun": true,
"waren": true,
"vill": true,
"koum": true,
"éischt": true,
"Januar": true,
"war,": true,
"No": true,
"gi": true,
"Juni": true,
"sinn.": true,
"Juli": true,
"zënter": true,
"beim": true,
"esou": true,
"ënnert": true,
"haut": true,
"Dezember": true,
"Eng": true,
"Oktober": true,
"September": true,
"ëmmer": true,
"August": true,
"ronn": true,
"lëtzebuergesche": true,
"Lëtzebuerger": true,
"nëmmen": true,
"gehéiert": true,
"géint": true,
"Mäerz": true,
"mee": true,
"puer": true,
"Loutrengen.": true,
"grouss": true,
"Februar": true,
"krut": true,
"kann": true,
"Spaweck.": true,
"Abrëll": true,
"hunn.": true,
"gëtt.": true,
"Dës": true,
"dräi": true,
"sou": true,
"Plaz": true,
"zur": true,
"Meter": true,
"Op": true,
"November": true,
"gouf,": true,
"éischte": true,
"Fir": true,
"nei": true,
"dës": true,
"nees": true,
"säin": true,
"geet": true,
"iwwert": true,
"u": true,
"Provënz": true,
"aner": true,
"Dat": true,
"war.": true,
"zwou": true,
"Lëtzebuerg,": true,
"Seng": true,
"him": true,
"dann": true,
"hire": true,
"deenen": true,
"Enn": true,
"konnt": true,
"Lëtzebuerg.": true,
"Jean": true,
"Kinnek": true,
"kënnt": true,
"sinn,": true,
"Kierch": true,
"Zu": true,
"Film": true,
"President": true,
"zwee": true,
"Frankräich": true,
"of": true,
"La": true,
"laang": true,
"nom": true,
"Zënter": true,
"Lëscht": true,
"duerno": true,
"steet": true,
"allem": true,
"bekannt": true,
"NGC": true,
"En": true,
"tëschent": true,
"A": true,
"Ufank": true,
"Joren": true,
"Bei": true,
"hunn,": true,
"Member": true,
"keng": true,
"selwer": true,
"hirem": true,
"Mat": true,
"Paräis": true,
"wann": true,
"m": true,
"Stär": true,
"Wéi": true,
"km": true,
"hat,": true,
"dacks": true,
"senge": true,
"hirer": true,
"Leit": true,
"deene": true,
"besteet": true,
"erëm": true,
"weider": true,
"dunn": true,
"Chr.\"": true,
"gemaach": true,
"d'": true,
"hiren": true,
"Duerch": true,
"franséische": true,
"dëser": true,
"Als": true,
"Dag": true,
"zesumme": true,
"Geschicht.": true,
"d'éischt": true,
"Louis": true,
"Héicht": true,
"goufe": true,
"Krich": true,
"Weltkrich": true,
"kleng": true,
"gebaut": true,
"Land": true,
"ouni": true,
"just": true,
"Awunner": true,
"Doud": true,
"Well": true,
"Säi": true,
"ongeféier": true,
"grousse": true,
"anere": true,
"besonnesch": true,
"Zweete": true,
"wéinst": true,
"Papp": true,
"Kéier": true,
"gemaach.": true,
"Paul": true,
"gegrënnt": true,
"genannt": true,
"leeft": true,
"Stärebild": true,
"nëmme": true,
"dësem": true,
"gutt": true,
"Charles": true,
"véier": true,
"hat.": true,
"haaptsächlech": true,
"Famill": true,
"D'": true,
"staark": true,
"Tour": true,
"Joseph": true,
"Och": true,
"gëtt,": true,
"à": true,
"von": true,
"meeschtens": true,
"genannt.": true,
"Pierre": true,
"di": true,
"Jong": true,
"Se": true,
"Le": true,
"Gare": true,
"Bis": true,
"schonns": true,
"sollt": true,
"dofir": true,
"goufen.": true,
"soll": true,
"lëtzebuergesch": true,
"éischter": true,
"zwéi": true,
"éischten": true,
"groussen": true,
"direkt": true,
"kënnen": true,
"Dall": true,
"kuerz": true,
"Nom": true,
"fréier": true,
"Dëst": true,
"also": true,
"spéider": true,
"z.B.": true,
"Nodeems": true,
"verschidde": true,
"und": true,
"Däitschland": true,
"och.": true,
"Kuckt": true,
"däitsche": true,
"bal": true,
"dëse": true,
"Liichtjoer": true,
"Fläch": true,
"drop": true,
"schonn": true,
"wäit": true,
"Robert": true,
"ware": true,
"Staat": true,
"laanscht": true,
"Form": true,
"Do": true,
"genannt,": true,
"da": true,
"zwar": true,
"Henri": true,
"al": true,
"hei": true,
"lëtzebuergeschen": true,
"Geschicht": true,
"ka": true,
"Beispill": true,
"Dem": true,
"nodeems": true,
"gréisste": true,
"I.": true,
"gréissten": true,
"däitsch": true,
"neie": true,
"spillt": true,
"Universitéit": true,
"während": true,
"Geographie.": true,
"John": true,
"gegrënnt.": true,
"héich": true,
"eréischt": true,
"belscher": true,
"sengen": true,
"New": true,
"waren,": true,
"elo": true,
"Europa": true,
"Laf": true,
}