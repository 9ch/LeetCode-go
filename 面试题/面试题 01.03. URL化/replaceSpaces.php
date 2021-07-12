<?php

class Solution
{

    /**
     * @param String $S
     * @param Integer $length
     * @return String
     */
    function replaceSpaces($S, $length)
    {
        $result = '';
        for ($i = 0; $i < $length; $i++) {
            if ($S[$i] == ' ') {
                $result += '%20';
            } else {
                $result += $S[$i];
            }
        }
        return $result;
    }
}