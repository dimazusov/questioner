### тз к вопросам(questioner) для программиста

разбиваем на части
составляем шаблон вопроса
поиск по шаблону вопроса ответа(sentence.Sentence)
замена найденной части
перемести {какое имя у файла или папки который нужно переименовать?}

извлекаем все вопросы
{какое имя у файла или папки который нужно переместить?} -> какое имя у файла или папки который нужно переместить?

необходимо выполнить mv {какое имя у файла или папки который нужно переместить?} {какое имя у файла или папки в которую нужно переместить?}

В бд хранятся:
 *какое имя у файла или папки который нужно переместить? -> [имя файла который нужно переместить]
 *какое имя у файла или папки который нужно переместить? -> [какое имя у файла или папки в которую нужно переместить]

 *1.txt - имя файла который нужно переместить
 *folder - имя папки в которуй нужно переместить

Для всех вопросов делаем:

1. перевод в нормальную форму(у первого имени существительного меняем падеж на именительный)
 поиск шаблона ответа по вопросу
 какое имя у файла или папки который нужно переименовать?
 ->
 имя файла который нужно переименовать
 имя папки которую нужно переименовать

2. составляем шаблон
 имя файла который нужно переименовать -> * - имя файла который нужно переименовать
 имя папки которую нужно переименовать -> * - имя файла который нужно переименовать

3. ищем по шаблону 
 {*} - имя файла который нужно переименовать -> 1.txt - имя файла который необходимо переименовать

4. проводим замену
 делим предложение по "-",
 изначальный вопрос заменяем на левую часть найденого предложения заменяем
 перемести {какое имя у файла или папки который нужно переименовать?} -> перемести 1.txt
