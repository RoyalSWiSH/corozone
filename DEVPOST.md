Wir versuchen mehrere Funktionen für den Nutzer in unserer App zu entwickeln. Für die einzelnen Aspekte wollen wir unbedingt mit mehreren der vielen anderen tollen Projekte kooperieren (Bluetooth Kontakt-tracking mit STRICT,  Intensivbetten auslastung, Supermarktauslastung, Psychologische Betreuung, Video Arzt...). Tested unseren Prototypen https://pr.to/W864FN/ (im Browser, oder auf dem Handy erst die proto.io App runterladen) oder schaut im (noch rudimentären) Git vorbei https://github.com/RoyalSWiSH/corozone/.

## Inspiration
Inspiration kam aus einem Tatendrang nicht in dieser Krise nur zuzuschauen zu müssen und der Chance unbürokatisch anzupacken wo Hilfe benötigt wird.

## Was macht es ?
Unsere am Wochenende implementierte App enthält aktuell die User Authentifizierung und kann Anfragen an ein Backend senden. An einem Backend für Nachbarschaftshilfe haben wir auch gearbeitet, aber beides noch nicht verknüpft.

Basierend auf unserem Prototypen soll die App folgendes können.
**Nachbarschaftshilfe**: Erstelle eine Anfrage für dich oder auch für deine Großeltern in einer anderen Stadt, deine Nachbarn nehmen sie an. Wenn man einmal einen Antrag von einer Person angenommen hat, versuchen wir auch weitere anfragen an diese Person weiterzureichen, um die Anzahl neue Kontakte zu reduzieren. Zeige außerdem den Menschen in deiner Region, welche Vorräte du auf Lager hast, du im Falle eines Falles teilen würdest, um Anreize für Hamsterkäufe zu vermeiden.

**Kontakt-Tracking**: Trage dein festes Cluster sozialer Kontakte ein und dann alle Ereignisse wo du außerhalb des Clusters Kontakte mit anderen hattest. Wir hoffen hier andere Projekte wie das STRICT Protokoll zu integrieren.

**Mini-Diagnose**: Trage deine Symptome ein und bekomme einen ersten Anhaltspunkt ob du dich bei einem Arzt melden solltest. Wenn ja, siehst du direkt, wann die Hotline überlastet ist und wenn du es erlaubst können deine Daten an deinen Arzt direkt weitergegeben werden, sodass dort Zeit gespart wird.

**Datenschutz**: Damit die App funktioniert müssen Daten geteilt werden. Aber du behälst die Kontrolle, da jede Anfrage auf deine Daten sichtbar in deinem Privacy Tab ist. Möchte ein Forscher zb Bewegungsprofile haben, um eine Publikation zu schreiben, bekommst du deine Anfrage, ob du deine Daten für diesen Zweck teilen möchtest.


## Wie habe ich es gebaut?

Das Frontend ist in nativescript vue implementiert, damit es sowohl nativ auf iOS wie auch auf Android läuft. Die User Authentifizierung läuft über Google Firebase, was später aber auch ersetzt werden kann. Im Backend setzen wir auf Go (Echo) und Java Spring.

## Herausforderungen

User Authentifizierung ist eine Wissenschaft für sich. Erst habe ich probiert das Rad in Go neu zu erfinden, bis mich der Mentor Dominik Guhr auf Google Firebase und Keycloak aufmerksam gemacht hat. Den Rest der Woche habe ich damit verbracht, verschiedene Tutorials durchzuarbeiten, die aber teilweise auf Nativescript 4.0 basierten, aktuell ist aber V6.5 und das musste zunächst migriert werden, bevor es lief. Da Nativescript nicht so verbreitet wie z.b. React ist, findet man auch nicht so viele Ressourcen.

## Stolz bin ich auf
Der Moment als "New user created" im Android Simulator erschienen ist und in der Firebase DB zu sehen war, da habe ich gejubelt. Auch wenn es als etwas sehr kleines erscheint, war das für mich ein großes Erfolgserlebnis.

## Was habe ich gelernt
Einigen Experten haben mir geholfen die DSGVO besser zu verstehen, Mentoren haben mir bei User Authentification geholfen, ich habe Feedback von professionellen UX-Designern bekommen und natürlich viel Erfahrung in nativescript 
und Authentifizierung mit Firebase gesammelt.

## Was kommt als nächstes für Corozone
Weiter entwickeln am nativescript frontend, Integration der Ideen aus den vielen anderen Projekten, testen der Nachbarschaftshilfe für Einkäufe mit meinen Nachbarn und Familie

* weiterentwickeln des Front-Ends in Nativescript Vue, API-Requests mit Axios implementieren
* Fertigstellung des Backends und integration ins Frontend
* Integration der Daten von Bett und Hotline Kapazitäten per Ampel mit Hilfe der Projekte anderer Teams
* sicheres Senden von Symptomen und persönlichen Daten an Ärzte
* integration von STRICT, das Projekt eines Teams um Kontakte über Bluetooth zu tracken
* middleware die alle requests auf persönliche Daten trackt und dem User im privacy Bereich zur Verfügung stellt
* detailliertere Ausarbeitung des Prototypen
* DSGVO-Konformität, evtl. mit offizieller Prüfung (da Gesundheitsdaten) vorbereiten
