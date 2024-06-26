package main

import (
	"flag"
	f "fmt"
	"strings"
)

func main() {
	var file_name string
	var java_package string
	var level string
	var help bool

	flag.StringVar(&file_name, "f", "common", "output file [common.log, error.log, rabbitmq.log]")
	flag.StringVar(&java_package, "p", "com.panbet", "package [com.worldline, com.rabbit]")
	flag.StringVar(&level, "l", "info", "level [warn, trace, info, error]")
	flag.BoolVar(&help, "h", false, "This page")
	flag.Parse()
	if help {
		flag.PrintDefaults()
	} else {
		logger_generator(java_package, file_name, level)
	}
}

func logger_generator(java_package string, file_name string, level string) {
	f.Print(`
<!-- START ` + strings.ToUpper(file_name) + `.log -->
<appender name="` + strings.ToUpper(file_name) + `_FILE" class="ch.qos.logback.core.rolling.RollingFileAppender">
    <file>${logPrefix}/` + file_name + `.log</file>
    <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
        <fileNamePattern>${logPrefix}/old/` + file_name + `.%d{"yyyy-MM-dd-HH" GMT}.log</fileNamePattern>
        <maxHistory>7</maxHistory>
    </rollingPolicy>
    <encoder>
        <pattern>%d{"yyyy-MM-dd-HH:mm:ss,SSS Z" GMT} [%t] [%c] %-6p%m%n</pattern>
    </encoder>
</appender>
<appender name="` + strings.ToUpper(file_name) + `_ROLL" class="ch.qos.logback.classic.AsyncAppender">
   <appender-ref ref="` + strings.ToUpper(file_name) + `_FILE"/>
</appender>
<logger name="` + java_package + `" additivity="false" level="` + strings.ToUpper(level) + `">
    <appender-ref ref="` + strings.ToUpper(file_name) + `_ROLL"/>
</logger>
<!-- END ` + strings.ToUpper(file_name) + `.log -->
`)
}
