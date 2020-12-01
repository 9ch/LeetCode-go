<?php

/**
 * mixed一群猴子排成一圈，按1，2，...，n依次编号。
 *
 * 然后从第1只开始数，数到第m只,把它踢出圈，
 *
 * 从它后面再开始数，再数到第m只，在把它踢出去...，
 *
 * 如此不停的进行下去，直到最后只剩下一只猴子为止，那只猴子就叫做大王。
 */

//维护一个链表、也可以成为数组。
function monkey($arr, $m = 3)
{
    $i = 1;

    while (count($arr) > 1) {
        if ($i % $m == 0) {
            unset($arr[$i - 1]);
        } else {
            array_push($arr, $arr[$i - 1]);
            unset($arr[$i - 1]);
        }
        $i++;
    }
    return array_pop($arr);
}


$arr = array('a', 'b', 'c', 'd', 'e', 'f', 'g', 'h');
var_dump(monkey($arr));
