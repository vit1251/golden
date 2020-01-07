package config

func Create(section string, name string, summary string) {

func Set(section string, name string, value string) {

}

func Get(section string, name string, defaultValue *string) (*string, error) {
}

func Init() {

	/* Step 1. Create "settings" store */
	sqlStmt := "CREATE TABLE IF NOT EXISTS settings (" +
		    "    settingSection CHAR(64) NOT NULL PRIMARY KEY," +
		    "    settingName CHAR(64) NOT NULL PRIMARY KEY," +
		    "    settingValue CHAR(64) NOT NULL PRIMARY KEY," +
		    "    UNIQUE(settingSection, settingName)"
		    ")"

	/* Step 2. Create "area" store */
	sqlStmt := "CREATE TABLE IF NOT EXISTS area (" +
		    "    areaId INTEGER NOT NULL PRIMARY KEY," +
		    "    areaName CHAR(64) NOT NULL," +
		    "    areaType CHAR(64) NOT NULL," +
		    "    areaPath CHAR(64) NOT NULL," +
		    "    areaSummary CHAR(64) NOT NULL," +
		    "    areaOrder INTEGER NOT NULL," +
		    "    UNIQUE(areaName)"
		    ")"

	/* Step 3. Create "rule" store (i.e. rules is way to implement carbon copy and etc.) */
//	sqlStmt := "CREATE TABLE IF NOT EXISTS rule (" +
//		    "    ruleId INTEGER NOT NULL PRIMARY KEY," +
//		    "    ruleName CHAR(64) NOT NULL" +
//		    ")"
//	sqlStmt := "CREATE TABLE IF NOT EXISTS ruleStep (" +
//		    "    ruleStepId INTEGER NOT NULL PRIMARY KEY," +     /* Rule step index                                               */
//		    "    ruleId INTEGER NOT NULL," +                     /* Rule step owner                                               */
//		    "    ruleStepType INTEGER NOT NULL," +               /* Type "Copy", "Move", "Delete" and etc.                        */
//		    "    ruleStepValue INTEGER NOT NULL," +              /* Comparision value                                             */
//		    ")"

	/* Step 3. Create settings */
	self.Create("main", "Origin", "Origin was provide BBS station name and network address")
	self.Create("main", "TearLine", "Tearline provide person sign in all their messages")
	self.Create("main", "TempInbound", "Directory where should be process incoming packets")
	self.Create("main", "TempOutbound", "Directory where process outbound packet")
	self.Create("main", "Address", "FidoNet network point address (i.e. POINT address)")
	self.Create("main", "Link", "FidoNet uplink provide (i.e. BOSS address)")

	/* Step 4. Initialize Golden common parameters */
	self.Set("main", "FirstName", "Alice")
	self.Set("main", "LastName", "Cooper")
	self.Set("main", "Country", "Russia")
	self.Set("main", "City", "Moscow")
	self.Set("main", "Origin", "Yo Adrian, I Did It! (c) Rocky II")
	self.Set("main", "TearLine", "Golden/LNX 1.2.1 2020-01-05 18:29:20 MSK (master)")
	self.Set("main", "Address", "2:5030/1592.15")
	self.Set("main", "Link", "2:5030/1592.0")
	self.Set("main", "Inbound", "/var/spool/ftn/inb")
	self.Set("main", "Outbound", "/var/spool/ftn/outb")

}
