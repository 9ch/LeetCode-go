<?php

/**
 * 某股票每日价格数据包括开盘价，最高价，最低价，收盘价，要求编写一个函数，通过输入一个多日的价格数据，
 * 输出每日收盘价的7日平均价（当日及前6日收盘价的平均价），例如：
 * 输入：
 * [[100,110,98,104],[104,110,100,108]]
 * 输出：
 * [[100,110,98,104，104],[104,110,100,108，106]]
 */

function gupiao($a)
{
    if (count($a) < 1) {
        return $a;
    }

    $temp = [];
    foreach ($a as $k => &$value) {
        if (count($temp) == 7) {
            array_shift($temp);
        }
        array_push($temp, $value[3]);
        $value[4] = round(array_sum($temp) / count($temp),2);
    }
    return $a;
}


$a = [
    [100, 110, 98, 104], [104, 110, 100, 108], [105, 110, 100, 110], [105, 110, 100, 112],
    [100, 110, 98, 104], [104, 110, 100, 108], [105, 110, 100, 110], [105, 110, 100, 112],
];
var_dump(gupiao($a));
