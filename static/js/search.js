// Load the search index from the JSON file
var searchIndex;
$.getJSON('/index.json', function(data) {
  searchIndex = lunr(function() {
    this.ref('url'); // Use the 'ID' field as the reference
    this.field('title');
    this.field('content');

    data.forEach(function(doc) {
      this.add(doc);
    }, this);
  });
});

// Perform a search when the user types in the search box
$('#searchInput').on('keyup', function() {
  var query = $(this).val();
  var results = searchIndex.search(query);

  var $searchResults = $('#searchResults');
  $searchResults.empty();

  // Display a message if no results were found
  if (results.length === 0) {
    $searchResults.append('<li>No results found</li>');
  } else {
    // Display the search results
    results.forEach(function(result) {
      console.log(result);
      var item = '<li><a href="' + result.ref + '">' + result.doc.title + '</a></li>';
      $searchResults.append(item);
    });
  }
});
