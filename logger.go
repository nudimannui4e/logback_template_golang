package main

import (
	f "fmt"
	"os"
)
func main() {

    logger := os.Args[1]
	logname := os.Args[2]
	level := os.Args[3]

	if len(os.Args[1:]) != 3 {
		usage()
	} else {
		logger_generator(logger, logname, level)
	}
}

func usage() {
	usage := `Скрипт использует 3 аргумента, для вывода logback.xml по шаблону
Пример: ./logger com.panbet.service.push.service.DeviceRegistrationService devices info
Где - 
  1.LOGGER - имя класса, пример: com.panbet.web.managers.account.AccountManager
  2.LOGNAME - просто имя файла лога, указывается без расширения(!)
  3.LEVEL - уровень логирования. info | debug
  `
  f.Print(usage)
}

func logger_generator(logger string, logname string, level string) {
	f.Print(`
<!-- START ` + logname + `.log -->
<appender name="` + logname + `_FILE" class="ch.qos.logback.core.rolling.RollingFileAppender">
    <file>${logPrefix}/` + logname + `.log</file>
    <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
        <fileNamePattern>${logPrefix}/old/` + logname + `.%d{"yyyy-MM-dd-HH" GMT}.log</fileNamePattern>
        <maxHistory>7</maxHistory>
    </rollingPolicy>
    <encoder>
        <pattern>%d{"yyyy-MM-dd-HH:mm:ss,SSS Z" GMT} [%t] [%c] %-6p%m%n</pattern>
    </encoder>
</appender>
<appender name="` + logname + `_ROLL" class="ch.qos.logback.classic.AsyncAppender">
   <appender-ref ref="` + logname + `_FILE"/>
</appender>
<logger name="` + logger + `" additivity="false" level="` + level + `">
    <appender-ref ref="` + logname + `_FILE"/>
</logger>
<!-- END ` + logname + `.log -->
	`)
}
