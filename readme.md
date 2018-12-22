Базовый пример можно посмотреть по этой ссылке  
https://vk.com/videos463580667?z=video-16108331_456259536%2Fpl_463580667_-2  

colly API reference  
https://godoc.org/github.com/gocolly/colly  

Селекторы  
GoQuery Selector is a selector used by https://github.com/PuerkitoBio/goquery  

goquery brings a syntax and a set of features similar to jQuery to the Go language.  

Селекторы jQuery можно посмотреть здесь:  
http://basicweb.ru/jquery/jquery_selectors.php  
http://jquery.page2page.ru/index.php5/%D0%A1%D0%B5%D0%BB%D0%B5%D0%BA%D1%82%D0%BE%D1%80%D1%8B  

ну или в яндексе  

В этом проекте попробуем дотянуться до сайта московской биржи и получить сведения об открытых позициях по деривативам,  
например  
https://www.moex.com/ru/contract.aspx?code=BR-1.19  

Учитываем, что код инструмента периодически меняется (BR-1.19 - каждый месяц)  

Также возможна ситуация, когда вместо реальной страницы с данными сайт сначала показывает лицензионное соглашение,  
где нужно нажать ссылку "Согласен". На самом деле страница уже готова, ЛС отображается поверх нее. Дополнение. На странице есть скрипт, проверяющий куки на предмет того, показывалось пользователю лицензионное соглашение или нет. Может быть получится генерировать кук, чтобы ЛС не отображалось.  

```
<script>
    var language = "ru";
    
    $(function () {

        if (typeof ($.cookie('disclaimerAgreementCookie')) != "undefined") {
            $("#dialog-confirm").hide();
            return;
        }

        //$("#content-disclaimer").load("/content/disclaimers/disclaimerru.html");
        $("#disclaimer-modal").load("/static-content/disclaimers/disclaimerru.html");
        $('.ui-dialog-buttonpane').hide();

        if (language == 'EN') $('.disclaimer__buttons a').eq(0).text('I Agree');
        if (language == 'EN') $('.disclaimer__buttons a').eq(1).text('I Do Not Agree');
        $('body').on('click', '.disclaimer__buttons a', function () {
            if ($(this).attr('data-dismiss') == 'modal') {
                $('#disclaimer-modal').fadeOut();
                $('.ui-widget-overlay').hide();
                $('.ui-dialog').hide();
                $.cookie("disclaimerAgreementCookie", 1, { expires: 365, path: '/' });
            }
        });


        $('#disclaimer-modal').show();

        if (typeof ($.cookie('disclaimerAgreementCookie')) != "undefined")
        {
            $("#dialog-confirm").hide();
            return;
        }
              
```


На нужной странице нас интересует таблица (скачайте этот файл, чтобы посмотреть исходный html-код):

 <input type="text" id="optDate" style="width: 70px;" ng-show="engine!='options' && isReadyContractSeries" />
    <strong ng-show="engine!='options' && isReadyContractSeries">Открытые позиции *</strong>
    <table class="contract-open-positions table1" ng-show="engine!='options' && isReadyContractSeries" style="margin-bottom: 0;">
        <tbody >
            <tr>
                <th rowspan="2"></th>
                <th colspan="2" class="white-border-column">Физические лица</th>
                <th colspan="2" class="white-border-column">Юридические лица</th>
                <th rowspan="2">Итого</th>
            </tr>
            <tr>
                <th>Длинные</th>
                <th>Короткие</th>
                <th>Длинные</th>
                <th>Короткие</th>
            </tr>
            <tr>
                <td>Открытые позиции</td>
                <td class="text_right">{{openOptions[0].PhysicalLong | limitTo:8 | number}}</td>
                <td class="text_right">{{openOptions[0].PhysicalShort | limitTo:8 | number}}</td>
                <td class="text_right">{{openOptions[0].JuridicalLong | limitTo:8 | number}}</td>
                <td class="text_right">{{openOptions[0].JuridicalShort | limitTo:8 | number}}</td>
                <td class="text_right">{{openOptions[0].Summary | limitTo:8 | number}}</td>
            </tr>
            <tr>
                <td>Изменение</td>
                <td class="text_right">{{openOptions[1].PhysicalLong | number}}</td>
                <td class="text_right">{{openOptions[1].PhysicalShort | number}}</td>
                <td class="text_right">{{openOptions[1].JuridicalLong | number}}</td>
                <td class="text_right">{{openOptions[1].JuridicalShort | number}}</td>
                <td class="text_right">{{openOptions[1].Summary | number}}</td>
            </tr>
            <tr>
                <td>Количество лиц</td>
                <td class="text_right">{{openOptions[2].PhysicalLong | number}}</td>
                <td class="text_right">{{openOptions[2].PhysicalShort | number}}</td>
                <td class="text_right">{{openOptions[2].JuridicalLong | number}}</td>
                <td class="text_right">{{openOptions[2].JuridicalShort | number}}</td>
                <td class="text_right">{{openOptions[2].Summary | number}}</td>
            </tr>
        </tbody>
    </table>
    <div ng-switch on="type" style="margin: 4px 0 16px;">
        <div ng-switch-when="F" style="text-align: left;">
            * Суммарно по всем фьючерсам, базовым активом которых является Сырая нефть сорта Brent
        </div>
        <div ng-switch-default  style="text-align: left;">
            * Суммарно по всем опционам , базовым активом которых является Сырая нефть сорта Brent, по всем страйкам
        </div>
    </div>

Обратите внимание, здесь применяется универсальный программный код для фьючерсов и опционов. Пока работаем с фьючерсами.  

Основной вопрос (пока) - как отработает шаблон при программном чтении страницы?:  

```
{{openOptions[0].PhysicalLong | limitTo:8 | number}}
```

Ответ - шаблон не отрабатывает :(

