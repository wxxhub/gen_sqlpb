package gen

import "testing"

func TestParseCreateTable(t *testing.T) {
	createTable := "CREATE TABLE `NewTable` (\n  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,\n  `name` varchar(32) NOT NULL DEFAULT '',\n  PRIMARY KEY (`id`)\n) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;"
	parseCreateTable(createTable)
}
