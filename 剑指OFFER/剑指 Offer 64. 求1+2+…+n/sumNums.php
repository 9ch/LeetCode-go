<?php

class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    public $ans = 0;
    function sumNums($n) {

        $this->ans += $n;
        $n > 0 && $this->sumNums($n-1);
        return $this->ans;

    }
}

$s = new Solution();
var_dump($s->sumNums(3));

//不使用类的方法
global $ans;
$ans = 0;
function sumNums($n) {

    $GLOBAL["ans"]+=$n;
    echo $n.' '.$GLOBAL["ans"].PHP_EOL;//不使用全局变量的话，会造成闭包问题。
    $n > 0 && sumNums($n-1);
    return $GLOBAL["ans"];
}

var_dump(sumNums(3));
