# News
___

Библиотека по работе с новостным API. Имеет два эндпоинта. Помогает получить 10 свежих новостей по категориям и по странам.

```shell
go get github.com/AlexCorn999/news
```
```go
...
  GetCategory(category Category) (*NewsResponse, error)
  GetNewsByCountry(category Category, country Country) (*NewsResponse, error)
...
```

# example 1

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AlexCorn999/news/news"
)

func main() {
	client, err := news.NewClient(time.Second * 5)
	if err != nil {
		log.Fatal(err)
	}

	articles, err := client.GetCategory(news.World)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(articles.Info())
}
```
```text
[Wed Nov  1 21:35:08 2023] GET https://gnews.io/api/v4/top-headlines?category=world&lang=ru&country=ru&max=5&apikey=88ea8ecf6571e60725e84639289b3ed1
Официальный Осло: Израиль "вышел за рамки соразмерности" при атаке на сектор Газа
Премьер-министр Норвегии Йонас Гар Стёре уверен, что израильские военные явно «вышли за рамки соразмерности» при атаке на сектор Газа, нарушая все правила ведения военных действий.
[IMAGE] https://img2.eadaily.com/r650x400/o/112/b5d8f7d1d66ee6d561351235d44a5.jpeg
[TIME] 2023-11-01T16:00:00Z
[URL] https://eadaily.com/ru/news/2023/11/01/oficialnyy-oslo-izrail-vyshel-za-ramki-sorazmernosti-pri-atake-na-sektor-gaza

Второй энергоблок сдан: "Росатом" завершил строительство Белорусской АЭС
Российская госкорпорация завершила строительство Белорусской АЭС. Второй энергоблок станции приняли в промышленную эксплуатацию.
[IMAGE] https://img8.eadaily.com/r650x400/o/f26/c8f8470344c6c521c6bf1fc6a2e50.jpeg
[TIME] 2023-11-01T13:59:00Z
[URL] https://eadaily.com/ru/news/2023/11/01/vtoroy-energoblok-sdan-rosatom-zavershil-stroitelstvo-belorusskoy-aes

Болгария выдворила российского журналиста по подозрению в шпионаже
Александр Гацак покинул Софию. Месяц он укрывался в российском посольстве
[IMAGE] https://gdb.rferl.org/0408CC96-EE8F-4E0B-ABEA-9C9B1F9E7AFB_cx0_cy6_cw0_w1200_r1.png
[TIME] 2023-11-01T13:56:49Z
[URL] https://www.svoboda.org/a/bolgariya-vydvorila-rossiyskogo-zhurnalista-po-podozreniyu-v-shpionazhe/32665958.html

Суд в Лондоне: РФ в споре о $60 млрд не может апеллировать к иммунитету
Речь идёт о споре с акционерами ЮКОСа
[IMAGE] https://gdb.rferl.org/058a0000-0aff-0242-ecb5-08dad47bbf63_w1200_r1.jpg
[TIME] 2023-11-01T12:55:08Z
[URL] https://www.svoboda.org/a/sud-v-londone-rf-v-spore-o-60-mlrd-ne-mozhet-apellirovatj-k-immunitetu/32665831.html

Украина сосредоточила 120-тысячную группировку на границе с Белоруссией - Молостов
Об этом заявил председатель белорусского государственного пограничного комитета полковник Константин Молостов, выступая в среду, 1 ноября в нижней палате парламента страны. 1 ноября 2023. EADaily
[IMAGE] https://img1.eadaily.com/r650x400/o/822/ee4fbe329b03e26a2e797c6dedbbd.jpeg
[TIME] 2023-11-01T12:41:00Z
[URL] https://eadaily.com/ru/news/2023/11/01/ukraina-sosredotochila-120-tysyachnuyu-gruppirovku-na-granice-s-belorussiey-molostov
```

```

# example 2

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AlexCorn999/news/news"
)

func main() {
	client, err := news.NewClient(time.Second * 5)
	if err != nil {
		log.Fatal(err)
	}

	articles, err := client.GetNewsByCountry(news.Science, news.Canada)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(articles.Info())
}
```
```text
[Wed Nov  1 21:36:44 2023] GET https://gnews.io/api/v4/top-headlines?category=science&lang=ru&country=ru&max=5&apikey=88ea8ecf6571e60725e84639289b3ed1
Заболеваемость COVID-19 в России за неделю выросла на 30%
Увеличение заболеваемости отмечено в 72 субъектах страны
[IMAGE] https://ysia.ru/wp-content/uploads/2023/03/596A9077-960x640-1.jpg
[TIME] 2023-11-01T05:34:14Z
[URL] https://ysia.ru/v-rossii-vyrosla-zabolevaemost-covid-19/

Очередной многокилометровый выброс пепла Ключевского сняли на видео
Расположенный на Камчатке вулкан Ключевская сопка продолжает извергаться, обновляя собственные рекорды. Так, в результате последних выбросов пепел поднялся на высоту до 14 километров. Об этом сообщили учёные Института вулканологии и сейсмологии ДВО РАН.
[IMAGE] https://static.life.ru/publications/2023/11/1/1069993497576.9155.jpg
[TIME] 2023-11-01T01:55:56Z
[URL] https://life.ru/p/1617970

В команде с роботами живые специалисты становятся ленивыми и невнимательными
Недавнее исследование указывает на то, что сотрудничество людей с роботами делает людей ленивее.
[IMAGE] https://news-ru.gismeteo.st/2023/01/shutterstock_1484047868.jpg
[TIME] 2023-10-31T19:35:11Z
[URL] https://www.gismeteo.ru/news/science/v-komande-s-robotami-zhivye-specialisty-stanovyatsya-lenivymi-i-nevnimatelnymi/

Засуха подарила ученым уникальное открытие у берегов Амазонки
 ‌Экстремально засушливая погода привела к снижению уровня Амазонки, и хотя это не совсем позитивная новость с точки зрения экологии, зато историкам повезло.
[IMAGE] https://news-ru.gismeteo.st/2023/07/shutterstock_1916509097.jpg
[TIME] 2023-10-31T19:20:35Z
[URL] https://www.gismeteo.ru/news/science/zasuha-podarila-uchenym-unikalnoe-otkrytie-u-beregov-amazonki/

NASA сообщило о проблемах с обработкой материалов с астероида Бенну
Первичный анализ веществ, обнаруженных вне основного бокса, позволил агентству сделать вывод: пробы включают и H2O, и углерод.
[IMAGE] https://news-ru.gismeteo.st/2023/10/20231011_zia_c218_031-scaled.jpeg
[TIME] 2023-10-31T19:14:01Z
[URL] https://www.gismeteo.ru/news/science/nasa-soobshhilo-o-problemah-s-obrabotkoj-materialov-s-asteroida-bennu/
```
___
