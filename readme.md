Базовый пример можно посмотреть по этой ссылке  
https://vk.com/videos463580667?z=video-16108331_456259536%2Fpl_463580667_-2

В этом проекте попробуем дотянуться до сайта московской биржи и получить сведения об открытых позициях по деривативам,  
например  
https://www.moex.com/ru/contract.aspx?code=BR-1.19  

Учитываем, что код инструмента периодически меняется (BR-1.19 - каждый месяц)  

Также возможна ситуация, когда вместо реальной страницы с данными сайт сначала показывает лицензионное соглашение,  
где нужно нажать ссылку "Согласен".  

На нужной странице нас интересует таблица:

'''
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
'''
Обратите внимание, здесь универсальный код для фьючерсов и опционов. Пока работаем с фьючерсами.  