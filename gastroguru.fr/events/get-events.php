<?php
function getEvents()
{
  $response = file_get_contents('http://localhost:44446/events');
  if ($response === false) {
    throw new Exception("Erreur lors de la récupération des événements");
  }

  $data = json_decode($response, true);
  if ($data === null || !isset($data['events'])) {
    throw new Exception("Erreur lors du décodage de la réponse JSON ou données manquantes");
  }

  return $data['events'];
}

try {
  $events = getEvents();

  foreach ($events as $event) {
    $dateString = $event['date'];
    $date = new DateTime($dateString);
    $day = $date->format('d');
    $monthName = $date->format('F');

    $eventHTML .= '
            <div class="col-md-4 animate-box">
                <div class="fh5co-event">
                    <div class="date text-center"><span>' . $day . '<br>' . $monthName . '</span></div>
                    <h3><a href="#">' . $event['title'] . '</a></h3>
                    <p>' . $event['description'] . '</p>
                    <p><a href="#">Read More</a></p>
                </div>
            </div>';
  }

  echo $eventHTML;
} catch (Exception $e) {
  echo "Une erreur s'est produite : " . $e->getMessage();
  return;
}

?>