������� ������ ����� ���������� �� ���� ������  
https://vk.com/videos463580667?z=video-16108331_456259536%2Fpl_463580667_-2

� ���� ������� ��������� ���������� �� ����� ���������� ����� � �������� �������� �� �������� �������� �� �����������,  
��������  
https://www.moex.com/ru/contract.aspx?code=BR-1.19  

���������, ��� ��� ����������� ������������ �������� (BR-1.19 - ������ �����)  

����� �������� ��������, ����� ������ �������� �������� � ������� ���� ������� ���������� ������������ ����������,  
��� ����� ������ ������ "��������".  

�� ������ �������� ��� ���������� �������:

'''
 <input type="text" id="optDate" style="width: 70px;" ng-show="engine!='options' && isReadyContractSeries" />
    <strong ng-show="engine!='options' && isReadyContractSeries">�������� ������� *</strong>
    <table class="contract-open-positions table1" ng-show="engine!='options' && isReadyContractSeries" style="margin-bottom: 0;">
        <tbody >
            <tr>
                <th rowspan="2"></th>
                <th colspan="2" class="white-border-column">���������� ����</th>
                <th colspan="2" class="white-border-column">����������� ����</th>
                <th rowspan="2">�����</th>
            </tr>
            <tr>
                <th>�������</th>
                <th>��������</th>
                <th>�������</th>
                <th>��������</th>
            </tr>
            <tr>
                <td>�������� �������</td>
                <td class="text_right">{{openOptions[0].PhysicalLong | limitTo:8 | number}}</td>
                <td class="text_right">{{openOptions[0].PhysicalShort | limitTo:8 | number}}</td>
                <td class="text_right">{{openOptions[0].JuridicalLong | limitTo:8 | number}}</td>
                <td class="text_right">{{openOptions[0].JuridicalShort | limitTo:8 | number}}</td>
                <td class="text_right">{{openOptions[0].Summary | limitTo:8 | number}}</td>
            </tr>
            <tr>
                <td>���������</td>
                <td class="text_right">{{openOptions[1].PhysicalLong | number}}</td>
                <td class="text_right">{{openOptions[1].PhysicalShort | number}}</td>
                <td class="text_right">{{openOptions[1].JuridicalLong | number}}</td>
                <td class="text_right">{{openOptions[1].JuridicalShort | number}}</td>
                <td class="text_right">{{openOptions[1].Summary | number}}</td>
            </tr>
            <tr>
                <td>���������� ���</td>
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
            * �������� �� ���� ���������, ������� ������� ������� �������� ����� ����� ����� Brent
        </div>
        <div ng-switch-default  style="text-align: left;">
            * �������� �� ���� �������� , ������� ������� ������� �������� ����� ����� ����� Brent, �� ���� ��������
        </div>
    </div>
'''
�������� ��������, ����� ������������� ��� ��� ��������� � ��������. ���� �������� � ����������.  