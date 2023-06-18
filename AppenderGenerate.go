package main

import (
	f "fmt"
	"flag"
)
func main() {
	var file_name string
	var java_package string
	var level string
//	var help string
/*	help := `Скрипт использует 3 аргумента, для вывода logback.xml по шаблону
Пример: ./logger com.panbet.service.push.service.DeviceRegistrationService devices info
Где - 
  1.LOGGER - имя класса, пример: com.panbet.web.managers.account.AccountManager
  2.LOGNAME - просто имя файла лога, указывается без расширения(!)
  3.LEVEL - уровень логирования. info | debug
  `
*/
	flag.StringVar(&file_name, "f", "common", "Default filename: common.log")
	flag.StringVar(&java_package, "j", "com.panbet", "Default package: com.panbet")
	flag.StringVar(&level, "l", "info", "Default logger level: info")
//	flag.TextVar(&help, "h", "man page", "If agrs < 3: help page")
	flag.Parse()
	logger_generator(java_package, file_name, level)
}

func logger_generator(java_package string, file_name string, level string) {
	f.Print(`
<!-- START ` + file_name + `.log -->
<appender name="` + file_name + `_FILE" class="ch.qos.logback.core.rolling.RollingFileAppender">
    <file>${logPrefix}/` + file_name + `.log</file>
    <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
        <fileNamePattern>${logPrefix}/old/` + file_name + `.%d{"yyyy-MM-dd-HH" GMT}.log</fileNamePattern>
        <maxHistory>7</maxHistory>
    </rollingPolicy>
    <encoder>
        <pattern>%d{"yyyy-MM-dd-HH:mm:ss,SSS Z" GMT} [%t] [%c] %-6p%m%n</pattern>
    </encoder>
</appender>
<appender name="` + file_name + `_ROLL" class="ch.qos.logback.classic.AsyncAppender">
   <appender-ref ref="` + file_name + `_FILE"/>
</appender>
<logger name="` + java_package + `" additivity="false" level="` + level + `">
    <appender-ref ref="` + file_name + `_FILE"/>
</logger>
<!-- END ` + file_name + `.log -->
`)
}
