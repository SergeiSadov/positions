### Запуск:
```sh
cp config.example.json config.json
sh build_ek.sh
sh deploy.sh
```
### Настройка логов:
Ждём, пока elasticserch станет доступен в kibana
Заходим на: http://localhost:5601/app/kibana#/management/kibana/index_pattern?_g=()
Добавляем паттерны:
error*
info*
### Деплой:
```sh
sh deploy.sh
```
### Мониторинг:
http://localhost:9099/targets
### Сервисы:
  - prometheus: localhost:9099
  - kibana: localhost:5601
  - elasticsearch: localhost:9020
  - pprof localhost:8000/debug/pprof
### Endpoints:
  - http://localhost:8000/positions
  - http://localhost:8000/summary