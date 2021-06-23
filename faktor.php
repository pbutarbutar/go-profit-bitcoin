<?php

function faktorial($n, $x) {
     $n = $n - 4;
     echo $n;
     echo "\n";
    if ($x > 1) {
        faktorial($n, $x-1);
    }
    
}

faktorial(128, 19);
