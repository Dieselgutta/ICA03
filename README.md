# ICA03


*Deltakere: Simen Fredriksen, Stian Blankenberg, Jone Manneråk, Kristian Hagberg, Tarjei Taxerås og Rasmus Sørby*


# Oppgave 1


Vi kjørte filen main.go i Git Bash. Da skriver vi ut en tabell med alle tegn i «string literal» ascii. VI får ut alle ascii-kode heksadesimalt med store bokstaver, symbol for ascii-kode og ascii kode binært. Slik at det blir seende slik ut: 


!Bilde1](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18554952_10158500684940411_717436644_n.png?oh=4c4b41d3b238065afffa539af3a49e54&oe=591D2D23)

Osv…


Vi fant ingen forskjeller når vi kjørte programmet på nettskyen (UH-IaaS).
På instansen i nettskyen (UH-IaaS) så utskriften slik ut:

![Bilde2](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18554792_10158500687760411_1203725809_n.png?oh=c457cd7a38510bbc7b241fa548a4a673&oe=591D8233)


Osv…




# 1b) 
Når vi lagde en funksjon greetingASCII() i samme filen ascii.go, som skriver ut "Hello :-)"  og utførte programmet på instansen i nettskyen (UH-IaaS) ble det seende slik ut:

![Bilde3](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18579106_10158500691750411_1873139703_n.png?oh=ed78dd88ca0a29676db739799965e575&oe=591CEE7A)


# c)
Koden for testen ligger i Github-repositoriet.





# Oppgave 2


# a)
Vi kjørte filen main.go til oppgave 2 i Git Bash og i nettskyen (UH-IaaS) . Da genererte den alle tallene og symbolene fra byteverdier, til Ascii. Det så slik ut(Git Bash til venstre og nettskyen til høyre):

![Bilde4](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18578525_10158500695125411_1586052045_n.png?oh=52b2986105eee18b89848142cccf3b8f&oe=591CF9E2)


Dette er bare en liten del av det som kom opp da vi kjørte programmet. Tankegangen vår var litt lik som i oppgave 1, men vi gjorde endringer for at den skulle iterere. Vi samarbeidet mye i gruppen for å komme fram til det riktige resultatet. Vi ser at vi får ett problem når vi skal printe ut ukjente symboler.

# b)
Tankegangen var lik som i oppgave 1. Vi brukte samme mainfunksjon som i oppgave a. Da vi kjørte programmet så det slik ut: Her ser vi at både oppgave a og b blir printet ut samtidig, dette er fordi vi har brukt samme main-fil til de to ettersom de hører til samme mappe (iso).

![Bilde5](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18516255_10158500697285411_698830966_n.png?oh=196ab55d7d0acdf17fd02a5df26f6c04&oe=591CF5C6)

Vi fikk ikke ut € symbolet, i steden viser nettskyen et rektangulært symbol forran “50”. 

# c)
Koden for testen ligger oppe på Github-repositoriet. 




# Oppgave 3

# a)
%s printer ut symboler for bytes av en string
 %q printer ut symboler for en tekststring 
%+q printer ut symboler for Unicode som er utenfor ASCII kodesett. 
%c printer ut symboler som den tolker innenfor utvidet ASII kodesett. 
I bytesekvensen må man endre slik at det blir «\x22\xC2\xBD\x3F\x3D\x3F\x20\xE2\x8C\x98\x22»
Vi får ikke ut ⌘ symbolet når vi printer, i steden viser den et rektangulært symbol.

![Bilde6](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18516165_10158500699905411_1881870183_n.png?oh=5aac8e6c63e54742c5f9b2b1b696558b&oe=591D7EDD)


# b)
Vi har under her skrevet ut lang01.wl, lang02.wl og lang03.wl. Resultatene vi fikk ser ut til å være de samme som var i de orginale filene, bare at øverst er de skrevet ut i heksadesimaler, og under er de skrevet ut i vanlig string. Men der ser vi også at noen av printene her ikke skrev de ut nøyaktig som de stod, dette er trolig fordi at de forskjellige språkene reagerer anderledes med den samme printen, som her ser ut til å ha en innstilling til å bruke det engelske kodesettet. Noe som resulterte i at alle tegnene som ikke er til felles med det engelske kodesettet ble seendes uleselig ut. Dette er selvfølgelig bare èn måte å vise hvilket kodesett orginalteksten ble skrevet med, andre metoder kan også være...

![Bilde7](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18516180_10158500703165411_271930854_n.png?oh=14e79df80eda4137269a560746047669&oe=591E3279)

![Bilde8](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18516211_10158500703740411_926807169_n.png?oh=0cf55844677b64b645e13916b351a03a&oe=591D2F6C)

![Bilde9](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18578754_10158500704000411_543887020_n.png?oh=428d75b1451760d1bf80402e64eb835c&oe=591D38A5)


“EF DA A3 D2 D3 CB 0A EF DA A3 D2 D3 CB C1 0A EF …” i unicode blir “u00ef, u00da, u00a3, u00d2, u00d3, u00cb, \n, u00ef, u00da, u00a3, u00d2, u00d3, u00cb, u00c1, \n, u00ef..”
 “FE FD 73 6B 61 72 0A FE FD 73 6B 61 72 61 6E 61 ...” i unicode blir “u00fe, u00fd, s, k, a, r, \n, u00fe, u00fd, s, k, a, r, n, a”.
“F8 79 65 73 70 65 73 69 61 6C 69 73 74 65 6E 0A ...” i unicode blir “  u00f8, y, e, s, p, e, s, i, a, l, i, s, t, e, n, \n”.




# c)

![Bilde10](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18492971_10158500708685411_1635529287_n.png?oh=42d014d1469bbd5c651b367f81e31e9a&oe=591CE887)





# Oppgave 4

# a) 
Ved å skrive jp bak filen får man Japanask 
Ved å skrive Is bak filen får man Islandsk

![Bilde11](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18578723_10158500708920411_765775719_n.png?oh=60e0c84d3a150c1cfad33ddb20bd3aea&oe=591D14AD)

# b) 
Ved å bruke U+23F0, får vi ut et klokke symbol, som vist i bildet.

![Bilde]12(https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18554429_10158500709035411_308355515_n.png?oh=d3bbb23131eb08da2da95590f02b7b36&oe=591CE464)
