# 指定された年月のカレンダーを出力するプログラムを作成してください。その際、日曜日は()で囲み、土曜日は[]で囲んで出力してください。
# !/usr/bin/env python3
# -*- coding: utf-8 -*-
import calendar
from datetime import *


# 日数を取得
def get_day_of_the_week(Year, Month):
    day_num = calendar.monthrange(Year, Month)
    return day_num[1]


if __name__ == "__main__":
    Year, Month = map(int, input().split('/'))
    # 閏年を考慮した日数
    day_num = get_day_of_the_week(Year, Month)
    # カレンダーを表示
    day = 1
    # カレンダーの文字列格納用
    calendar_line = ''
    for i in range(day_num):
        # yyyy/MM/ddのフォーマットから曜日ごとの数字を返す(0:月曜、6:日曜)
        check_day = datetime.strptime(str(Year) + '/' + str(Month) + '/' + str(day), '%Y/%m/%d')
        week_day = check_day.weekday()
        # 月初の曜日に応じてスペースを空ける
        if day == 1:
            space = ' ' * (week_day * 2)
            four_space = '    '
            if week_day == 6:
                calendar_line = calendar_line + '(' + str(day) + ')'
            elif week_day == 5:
                calendar_line = calendar_line + four_space + space + '[' + str(day) + ']'
                print(calendar_line + '\n')
                calendar_line = ''
            else:
                calendar_line = four_space + calendar_line + ' ' + space + str(day)
        # 土曜日は改行を含めた出力後初期化する
        elif week_day == 5:
            calendar_line = calendar_line + '[' + str(day) + ']'
            print(calendar_line + '\n')
            calendar_line = ''
        # 日曜日
        elif week_day == 6:
            calendar_line = calendar_line + '(' + str(day) + ')'
        elif day < 10:
            calendar_line = calendar_line + ' ' + str(day)
        elif day >= 10:
            calendar_line = calendar_line + str(day)
        if day >= day_num:
            print(calendar_line)
        day += 1
