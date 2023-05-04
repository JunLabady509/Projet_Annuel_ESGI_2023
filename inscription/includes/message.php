<?php 
if(isset($_GET['message']) && !empty($_GET['message']) && isset($_GET['type'])){
    echo '<div class="alert alert-' . $_GET['type'] . '" role="alert">' . htmlspecialchars($_GET['message']) . '</div>';
}
?>