<?php

class Solution
{

    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    function subsets($nums)
    {
        $length = count($nums);
        $result = [[]];
        for ($i = 0; $i < $length; $i++) {
            $temp = $result;
            for ($j = 0; $j < count($temp); $j++) {
                array_push($temp[$j], $nums[$i]);
            }
            $result = array_merge($result,$temp);
        }
        return $result;
    }
}