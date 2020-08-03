package web

// Code generated on "Mon Aug  3 14:01:58 PDT 2020" DO NOT EDIT.

const indexSource = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>CertGraph</title>
<link href="https://maxcdn.bootstrapcdn.com/bootswatch/3.3.7/cerulean/bootstrap.min.css" rel="stylesheet" integrity="sha384-zF4BRsG/fLiTGfR9QL82DrilZxrwgY/+du4p/c7J72zZj+FLYq4zY00RylP9ZjiT" crossorigin="anonymous">
<script src="https://d3js.org/d3.v4.min.js"></script>
<style type="text/css">
.links line {
  stroke-opacity: 0.6;
  stroke-width: 1px;
  fill: none;
}
.nodes circle {
  stroke: #333;
  stroke-width: 1.5px;
}
.upload-drop-zone {
  height: 200px;
  border-width: 2px;
  margin-bottom: 20px;
  color: #ccc;
  border-style: dashed;
  border-color: #ccc;
  line-height: 200px;
  text-align: center
}
.upload-drop-zone.drop {
  color: #222;
  border-color: #222;
}
</style>
</head>
<body>
<div class="container">

<nav class="navbar navbar-inverse">
    <div class="container-fluid">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">CertGraph</a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
            <ul class="nav navbar-nav">
                <!-- <li><a href="#">Graph</a></li> -->
            </ul>
            <ul class="nav navbar-nav navbar-right">
                 <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">Data <span class="caret"></span></a>
          <ul class="dropdown-menu" role="menu">
            <li><a href="#" data-toggle="modal" data-target="#URLmodal">URL</a></li>
            <li><a href="#" data-toggle="modal" data-target="#Pastemodal">Paste</a></li>
            <li><a href="#" data-toggle="modal" data-target="#Filemodal">File Upload</a></li>
          </ul>
        </li>
            </ul>
        </div>
    </div>
</nav>

<div class="panel panel-info">
  <div class="panel-heading">
    <h3 class="panel-title pull-left">Graph</h3>
    <div class="pull-right"><a href="#" class="btn btn-primary btn-xs" id="generate">Download SVG</a></div>
    <div class="clearfix"></div>
  </div>
    <svg id="graph" width="100%" height="500"></svg>
</div>

<div class="panel panel-info">
  <div class="panel-heading">Info</div>
  <div class="panel-body" id="node-info">
  </div>
</div>

<ul class="nav nav-tabs">
  <li class="active"><a href="#domains" data-toggle="tab" aria-expanded="false">Domains</a></li>
  <li class=""><a href="#certificates" data-toggle="tab" aria-expanded="true">Certificates</a></li>
</ul>
<div id="myTabContent" class="tab-content panel-body">
  <div class="tab-pane fade active in" id="domains">
    
<table class="table table-striped table-hover ">
  <thead>
    <tr>
      <th>#</th>
      <th>Domain</th>
      <th>Status</th>
      <th>Lookup</th>
    </tr>
  </thead>
  <tbody id="domain-list">
  </tbody>
</table> 

  </div>
  <div class="tab-pane fade" id="certificates">
    
<table class="table table-striped table-hover ">
  <thead>
    <tr>
      <th>#</th>
      <th>Hash</th>
      <th>Lookup</th>
    </tr>
  </thead>
  <tbody id="cert-list">
  </tbody>
</table> 

  </div>

</div>

<footer>
<hr>
<div class="row">
    <div class="col-xs-10"><a href="https://github.com/lanrat/certgraph">CertGraph</a></div>
</div>
</footer>

<!-- URL Modal -->
<div class="modal fade" id="URLmodal" role="dialog">
  <div class="modal-dialog">
    <!-- Modal content-->
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal">&times;</button>
        <h4 class="modal-title">Data URL</h4>
      </div>
      <div class="modal-body">
        <label for="inputURL" class="col-lg-2 control-label">URL</label>
        <div class="col-lg-10">
          <input type="text" class="form-control" id="inputURL" placeholder="https://domain.com/data.json">
        </div>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-primary" data-dismiss="modal" id="loadURL">Load</button>
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
      </div>
    </div>
  </div>
</div> <!-- /URL Modal -->

<!-- Paste Modal -->
<div class="modal fade" id="Pastemodal" role="dialog">
  <div class="modal-dialog">
    <!-- Modal content-->
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal">&times;</button>
        <h4 class="modal-title">JSON Data</h4>
      </div>
      <div class="modal-body">
        <label for="inputPaste" class="col-lg-2 control-label">URL</label>
        <div class="col-lg-10">
          <textarea class="form-control" rows="10" id="inputPaste"></textarea>
        </div>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-primary" data-dismiss="modal" id="loadPaste">Load</button>
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
      </div>
    </div>
  </div>
</div> <!-- /Paste Modal -->

<!-- File Modal -->
<div class="modal fade" id="Filemodal" role="dialog">
  <div class="modal-dialog">
    <!-- Modal content-->
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal">&times;</button>
        <h4 class="modal-title">JSON File</h4>
      </div>
      <div class="modal-body">
        <label for="inputFile" class="col-lg-2 control-label">File</label>
        <div class="col-lg-10">
         <input type="file" class="form-control file" id="inputFile">
        </div>
        <br/>
        <label for="drop-zone" class="col-lg-2 control-label">Or drag and drop a file below</label>
        <div class="upload-drop-zone" id="drop-zone">
          Just drag and drop a JSON file here
        </div>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-primary" data-dismiss="modal" id="loadFile">Load</button>
        <button type="button" class="btn btn-default" data-dismiss="modal" id="fileClose">Close</button>
      </div>
    </div>
  </div>
</div> <!-- /File Modal -->

</div> <!-- /container-->
<script src="//code.jquery.com/jquery-1.11.3.min.js"></script>
<script src="//maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
<script src="//cdn.rawgit.com/eligrey/FileSaver.js/e9d941381475b5df8b7d7691013401e171014e89/FileSaver.min.js"></script>
<script>
var svg = d3.select("svg");
var width = window.innerWidth-100;
//var width = svg.attr("width");
var height = svg.attr("height");

/*var svgElem = document.getElementById("graph");
var width = svgElem.width.animVal.value,
    height = svgElem.height.animVal.value;
console.log(width, height);*/

// TODO THIS
// http://www.coppelia.io/2014/07/an-a-to-z-of-extra-features-for-the-d3-force-layout/

var color = d3.scaleOrdinal(d3.schemeCategory10);
var simulation;

svg = svg.call(d3.zoom().on("zoom", zoomed)).append("g");

svg.append("defs").append("marker")
  .attr("id", "arrow")
  .attr("viewBox", "0 -5 10 10")
  .attr("refX", 20)
  .attr("refY", 0)
  .attr("markerWidth", 8)
  .attr("markerHeight", 8)
  .attr("orient", "auto")
  //.attr("stroke", function(d) { return color(d.type); })
  .append("svg:path")
  .attr("d", "M0,-5L10,0L0,5");

function resetGraph() {
  d3.select("g").selectAll("*").remove();
  createTables();

  // reset info
  var el = document.getElementById("node-info");
  el.innerText = "Click on a node in the graph to view details.";

  // redo layout
  simulation = d3.forceSimulation()
    .force("link", d3.forceLink().id(function(d) { return d.id; }))
    .force("charge", d3.forceManyBody().strength(-100))
    .force("center", d3.forceCenter(width / 2, height / 2));
}

function createGraph (error, graph) {
  if (error) throw error;

  var link = svg.append("g")
      .attr("class", "links")
    .selectAll("line")
    .data(graph.links)
    .enter().append("line")
      .attr("stroke", function(d) { return color(d.type); })
       .attr("marker-end", "url(#arrow)");

  var text = svg.append("g").attr("class", "labels").selectAll("g")
    .data(graph.nodes)
  .enter().append("g");

  text.append("text")
    .attr("x", 14)
    .attr("y", ".31em")
    .style("font-family", "sans-serif")
    .style("font-size", "0.7em")
    .text(function(d) { if (d.type == "domain") {return d.id; } return d.id.substring(0,8); });

  var node = svg.append("g")
      .attr("class", "nodes")
    .selectAll("circle")
    .data(graph.nodes)
    .enter().append("circle")
      .attr("r", 10)
      .attr("fill", function(d) { if (d.root == "true") return color(d.root); return color(d.type); })
      .call(d3.drag()
          .on("start", dragstarted)
          .on("drag", dragged)
          .on("end", dragended));

  node.on("click",function(d){
    // console.log("clicked", d.id);
    // console.log(d);
    updateInfoBox(d);
  });

  node.append("title")
      .text(function(d) { return d.id; });

  simulation
      .nodes(graph.nodes)
      .on("tick", ticked);

  simulation.force("link")
      .links(graph.links);

  function ticked() {
    link
        .attr("x1", function(d) { return d.source.x; })
        .attr("y1", function(d) { return d.source.y; })
        .attr("x2", function(d) { return d.target.x; })
        .attr("y2", function(d) { return d.target.y; });

    node
        .attr("cx", function(d) { return d.x; })
        .attr("cy", function(d) { return d.y; });
    text
        .attr("transform", function(d) { return "translate(" + d.x + "," + d.y + ")"});
  }
  createTables();
}


function dragstarted(d) {
  if (!d3.event.active) simulation.alphaTarget(0.3).restart();
  d.fx = d.x;
  d.fy = d.y;
}

function dragged(d) {
  d.fx = d3.event.x;
  d.fy = d3.event.y;
}

function dragended(d) {
  if (!d3.event.active) simulation.alphaTarget(0);
  d.fx = null;
  d.fy = null;
}

function zoomed() {
  svg.attr("transform", "translate(" + d3.event.transform.x + "," + d3.event.transform.y + ")" + " scale(" + d3.event.transform.k + ")");
}

d3.select("#generate").on("click", writeDownloadLink);
function writeDownloadLink(){
    try {
        var isFileSaverSupported = !!new Blob();
    } catch (e) {
        alert("blob not supported");
    }

    var html = d3.select("svg")
        .attr("title", "graph") //TODO
        .attr("version", 1.1)
        .attr("xmlns", "http://www.w3.org/2000/svg")
        .node().outerHTML;

    var blob = new Blob([html], {type: "image/svg+xml"});
    saveAs(blob, "certificate_graph.svg"); //TODO root node name
};

function updateInfoBox(d) {
  if (d) {
    var el = document.getElementById("node-info");
    var s = "Type: "+d.type+"</br>";
    if (d.type == "domain") {
      s = s + "Domain: "+linkifyDomain(d)+"</br>";
      s = s + "Status: "+d.status+"</br>";
    }else if (d.type = "certificate") {
      s = s + "Hash: "+linkifyCert(d)+"</br>";
    }
    el.innerHTML = s;
  }
}

function createTables() {
  // TODO: redo this in native d3
  domainEl = document.getElementById("domain-list");
  domain_tbody2 = document.createElement('tbody');
  domain_tbody2.id="domain-list";
  domainEl.parentNode.replaceChild(domain_tbody2, domainEl);

  certEl = document.getElementById("cert-list");
  cert_tbody2 = document.createElement('tbody');
  cert_tbody2.id="cert-list";
  certEl.parentNode.replaceChild(cert_tbody2, certEl);

  var domainCount = 0;
  function addTableDomain(d) {
    //console.log("domain", d);
    var c = "";
    if (d.root == "true") {
      c = "info";
    }
    $('#domain-list').append('<tr class="'+c+'"><td>'+ ++domainCount +'</td><td>'+linkifyDomain(d)+'</td><td>'+d.status+'</td><td>'+linkifyAny(d)+'</td></tr>');
  }

  var certCount = 0;
  function addTableCert(d) {
    //console.log("cert", d);
    $('#cert-list').append('<tr><td>'+ ++certCount + '</td><td>'+linkifyCert(d)+'</td><td>'+linkifyAny(d)+'</td></tr>');
  }


  d3.selectAll('circle').each(function(d){
    if (d.type == "domain") {
      addTableDomain(d);
    }else if (d.type == "certificate") {
      addTableCert(d);
    } else {
      console.log("Unknown Type: ", d.type);
    }
  })
}

function linkifyCert(d) {
  return '<a target="_blank" href="https://crt.sh/?sha256='+d.id+'">'+d.id+'</a>';
}
function linkifyDomain(d) {
  return '<a target="_blank" href="https://'+d.id+'">'+d.id+'</a>';
}
function linkifyAny(d) {
  return '<a target="_blank" href="https://crt.sh/?q='+d.id+'">&#x1F50E;</a>';
}

function getQueryVariable(variable){
  var query = window.location.search.substring(1);
  var vars = query.split("&");
  for (var i=0;i<vars.length;i++) {
    var pair = vars[i].split("=");
    if(pair[0] == variable){return pair[1];}
  }
  return "";
}

// new data from url
d3.select("#loadURL").on("click", loadURL);
function loadURL(){
  var url = document.getElementById("inputURL").value;
  history.pushState('', 'CertGraph', "?data="+url);
  resetGraph();
  d3.json(url, createGraph); 
}

// new data from paste
d3.select("#loadPaste").on("click", loadPaste);
function loadPaste(){
  var dataStr = document.getElementById("inputPaste").value;
  history.pushState('', 'CertGraph', "?");
  resetGraph();
  var data = JSON.parse(dataStr);
  createGraph(null, data);
}

// new data from paste
d3.select("#loadFile").on("click", loadFile);
function loadFile(){
  var file = document.getElementById("inputFile").files[0];
  history.pushState('', 'CertGraph', "?");
  resetGraph();
  var reader = new FileReader();
  reader.onload = function(e) {
    var dataStr = reader.result;
    var data = JSON.parse(dataStr);
    createGraph(null, data);
  }
  reader.readAsText(file);
}

var dropbox = document.getElementById('drop-zone');
function dragenter(e) {
  e.stopPropagation();
  e.preventDefault();
  dropbox.className = 'upload-drop-zone drop';
  // console.log("enter");
  return false;
}
function dragover(e) {
  e.stopPropagation();
  e.preventDefault();
  // console.log("over");
}
function dragleave(e) {
  e.stopPropagation();
  e.preventDefault();
  dropbox.className = 'upload-drop-zone';
  // console.log("leave");
  return false;
}
function drop(e) {
  // console.log("drop");
  e.stopPropagation();
  e.preventDefault();
  dropbox.className = 'upload-drop-zone';

  var dt = e.dataTransfer;
  var files = dt.files;

  var reader = new FileReader();
  reader.onload = function(e) {
    var dataStr = reader.result;
    var data = JSON.parse(dataStr);
    resetGraph();
    createGraph(null, data);
  }
  reader.readAsText(files[0]);
  $('#fileClose').click();
  return false;
}
dropbox.addEventListener("dragenter", dragenter, false);
dropbox.addEventListener("dragover", dragover, false);
dropbox.addEventListener("drop", drop, false);
dropbox.addEventListener("dragleave",dragleave, false);

// load initial graph data
var dataURL = getQueryVariable("data");
if (dataURL == "") {
  // default graph
  dataURL = "https://gist.githubusercontent.com/lanrat/8187d01793bf3e578d76495182654206/raw/c49741b5206d81935febdf563452cc4346381e52/eff.json";
}
resetGraph();
d3.json(dataURL, createGraph); 
</script>
</body>
</html>`

