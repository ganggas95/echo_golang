
var fileCollection = new Array();
var formAddProduct = $('#addProduct');
function readURL(input) {
    
    var reader = new FileReader();
    reader.onload = function (e) {
        $('img#upload-preview').attr('src', e.target.result);
    }
    reader.readAsDataURL(input.files[0]);
}
$("#singleImage").change(function(){
    readURL(this);
});

$('#images').change(function(e){
    var files= e.target.files;
    for(var i=0;i<files.length;i++)
    {
        fileCollection.push(files[i]);
        $('#image_preview')
            .append("<div class='col-md-3'><img class='img-responsive img-thumbnail' src='"+URL.createObjectURL(e.target.files[i])+"'></div>");
    }

});

formAddProduct.on('submit', function(e){
    e.preventDefault();
    var token = $(this).find('#_token').attr('value');
    var name_product = $(this).find('#name_product').attr('value');
    var price = $(this).find('#price').attr('value');
    var desc = $(this).find('#desc').attr('value');
    var category = $(this).find('#category').attr('value');
    console.log(desc);      
    /*var formData = new FormData($(this));
    for (var i = 0;i<fileCollection.length;i++){
        formData.append('images', fileCollection[i]);
    }
    formData.append('_token', token);
    formData.append('name_product', name_product);
    formData.append('price', price);
    formData.append('desc', desc);
    formData.append('category', category);
    var request = new XMLHttpRequest();
    request.open('post', '/admin/stock/add', true);

    request.send(formData);
    */
  
});
