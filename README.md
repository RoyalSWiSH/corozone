#supermarktstatus #krankenhausstatus #hotlinestatus #nachbarschaftshilfe #mini-diagnose #kontakt-tracking 

### Was ist die Herausforderung / das Problem? Was ist das Bedürfnis dahinter?
Beispiel: Problem: Häufung von Anrufen, Dringlichkeit | Bedürfnis: Effiziente Informationsvermittlung an Bürger:innen

Die App soll eine integrierte Lösung werden als Anlaufstelle für Bürger und regionale Nachbarschaftshilfe (Einkäufe, Gassi gehen, Kinder betreuen) und Entlastung für Krankenhäuser und Gesundheitsämter, indem Menschen gezielt zur richtigen Zeit an die richtige Adresse geführt werden, abhängig von der schwere der Symptome und der Auslastung der Stellen, Krankenhäuser und Hotlines.

Die Challenge ist, dass es trotz gut gemeinter Hilfe nicht zu mehr Kontakten kommt (#flattenthekurve) und dies im Einklang mit unseren Werten und dem Datenschutz zu gestalten. Wir wollen keine Überwachungsstaat, trotzdem wollen wir Daten nutzen. Ein Blanko-Datenschutzfreibrief per AGB ist aber nicht der richtige Weg. Wir erreichen das, indem wir dem Nutzer Kontrolle über seine Daten geben. Leiten wir Informationen über Einkäufe an einen anderen Nutzer weiter? Schau im Privacy Bereich nach, wann, warum und an wen wir die Daten weitergeleitet haben. Arten dir die Anzahl der Anfragen aus? Schalte das teilen der Daten in Zukunft einfach ab. Möchtest du deine Symptome direkt an die Notfallhotline weiterleiten, bevor du anrufst? Kein Problem, aber nur mit deinem Einverständnis, ansonsten bleiben die Daten verschlüsselt nur für dich zugänglich oder gleich einfach nur in deiner Tasche und nicht in der Cloud. Zugriff bekommt nur der, dem du deine Daten aktiv zeigst - in der digitalen oder in der realen Welt. Bemühung allein reicht aber nicht, damit wir das auch garantieren können müsste ein Pentester die App auf ihre Sicherheit prüfen.

Dazu müssen wir mit Organisationen wie der WHO oder dem RKI kooperieren. Wir sind keine Mediziner, aber stellen eine Mini-Diagnose entsprechend der Richtlinien des RKI zur Verfügung. Ob wir richtig gearbeitet haben, muss also bereits in der Entwicklungsphase ein Mediziner, Epidemiologe und/oder Virologe bewerten und gegebenenfalls Rückmeldung und Vorschläge geben. Wir wollen kein Aktionismus, sondern Hilfeleistende so gut es geht unterstützen.
https://github.com/RoyalSWiSH/corozone
### 2. Wer ist betroffen? Wem soll die Lösung konkret helfen?
Beispiele: Mitarbeiter:innen im Gesundheitsamt, Alleinerziehende, die im Home-Office arbeiten müssen

Bürger, die Einkäufe erledigen und Nachbarn helfen wollen.
Gesundheitsämter, indem Nutzer vorher wissen wie ausgelastet die Hotline ist und von selbst erst später anrufen, im Idealfall kann man grundsätzliche Informationen direkt weiterleiten bevor der Nutzer anruft und so Zeit sparen.
Krankenhäuser, indem Bürger, die ganz klar keine Sauerstoffversorgung im Krankenhaus benötigen, direkt zu eingerichteten emergency sites (umgewandelte Hotels...) oder Ärzten geleitet werden, die sie ebenfalls ausreichend betreuen können.
###3. Was ist der Themenbereich?
Kommunikation & Informationsvermittlung an Bürger*innen
Medizinische Versorgung
Prävention von Verbreitung des COVID-19
Solidarität und Zusammenhalt

###4. Wie würdest du die Herausforderung formulieren?
Beispiel: Wie können wir die große Zahl an Anrufern bewältigen, um sicher zu stellen, dass jede:r Bürger:in in max. 2h die benötigten Informationen erhält?

Die primäre Challenge für den Hackathon ist ein Frontend zu entwickeln z.b. mit vue native, Flutter, SwiftUI oder kotlin...das mit dem Server per JSON API kommuniziert. Die API, Login, Middleware ... werden in go implementiert, kann auch dabei Hilfe gebrauchen. Das ganze muss dann auch skalieren z.B. mit einen oder mehreren readonly backups des Datenbankserver und anderen best practices. Wichtig sind auch UI/UX Designer, die z.b. den Prototypen vor potentiellen Nutzern und den Stakeholdern präsentieren und Feedback bekommen, damit Entwickler nicht viel Zeit mit am ende nutzlosen Features verbringen.

5. Gibt es bereits Lösungsansätze oder internationale Vorbilder für den Umgang mit der Herausforderung?
Lösungsansätze, Ideen, Links bitte hier einfügen

Es gibt einen Prototypen als Diskussionsgrundlage, ein Datenbankschema für eine postgreSQL Datenbank und ein Backend in Go ist in Entwicklung. Kommuniziert wird zb im Slack des Hackathon oder per Zoom Videokonferenz, gearbeitet wird per Git. Jeder kann dort Issues für Features erstellen und jeder kann Issues annehmen.

Prototyp: https://pr.to/W864FN/

### Als einreichende Person musst du eine der folgenden Rollen übernehmen:
Koordination
Pat:in
inhaltliche Beratung
Hacker:in

###7 Was sind die wichtigsten Stakeholder?
bitte so präzise wie möglich

Das RKI (oder WHO) muss die App vorab prüfen. Apple gibt nur Apps mit Unterstützung von Internationalen Organisationen frei, um Falschinformationen zu vermeiden. Wir wollen eine Prüfung von Experten nicht umgehen, sondern darauf eingehen. Krankenhäuser und Hotlines müssen Informationen über ihre Auslastung zur verfügung stellen.
