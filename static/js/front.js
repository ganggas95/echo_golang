var fileCollection = new Array();
		var token = $('#_token').attr('value');

		$(document).on('change','#images',function(e){
			e.preventDefault();
			var files  = e.target.files;
			
			$.each(files, function(k, file){
				var reader = new FileReader();
				reader.readAsDataURL(file);
				reader.onload = function(e){
					var template = '<form name="uploadImage'+k+'">'+
								'<div class="col-md-3">'+
								'<input type="hidden" name="_token" value="'+token+'">'+
								'<div class="hovereffect">'+
							    	'<img class="img-responsive" src="'+e.target.result+'">'+
									'<div class="overlay">'+
									    '<a class="info">Set primary</a>'+
									    '<a class="info remove" role="button">Delete</a>'+
									'</div>'+
								'</div>'+
								'</div>'+
							'</form>';
						
					$('#image-preview').append(template);
					$('.remove').click(function(){
						console.log($(this).parent());
					})
				}; 
			});
		});

		$(document).on('load', '#images', function(e){
			e.preventDefault();
			var files  = e.target.files;
			$.each(files, function(k, file){
				fileCollection.push(file);

			});

		});
		$('#saveProduct').on('click', function(e){
			e.preventDefault();
			saveProduct();
		});



		function saveProduct(){
			var myForm = $('form[name="formAddProduct"]');
			var token = myForm.find('input[name="_token"]').val()
			var name_product = myForm.find('input[name="name_product"]').val();
			var price = myForm.find('input[name="price"]').val();
			var category = myForm.find('select[name="category"]').val();
			var desc = myForm.find('textarea[name="desc"]').val();
			var dataSend = {
				_token : token,
				product_name: name_product,
				product_price : price,
				product_category : category,
				product_desc : desc
			}

			$.post('/admin/stock/add', dataSend, function(data){
				productImages(data);
			});
			return false;
		}

		function productImages(id){
			for (var k = 0;k<fileCollection.length;k++){
				var myForm = $('form[name="uploadImage'+k+'"')[0];
				var formData = new FormData(myForm);
				formData.append('images',fileCollection[k]);
				formData.append('product_id', id);
				$.ajax({
				    url: '/admin/stock/image/upload',
				    data: formData,
				    cache: false,
				    processData: false,
				    contentType:false,
				    type: 'POST',
				    success: function ( data ) {
				        console.log(data);
				    }
				});
			}
		}



		var selDiv = "";
	    var storedFiles = [];
	    
	    $(document).ready(function() {
	        $("#files").on("change", handleFileSelect);
	        
	        selDiv = $("#image-preview"); 
	        //$("form[name='formAddProduct'").on("submit", handleForm);
	        
	        $("body").on("click", ".selFile", removeFile);
	    });

	    function handleFileSelect(e) {
	        var files = e.target.files;
	        var filesArr = Array.prototype.slice.call(files);
	        filesArr.forEach(function(f) {          

	            if(!f.type.match("image.*")) {
	                return;
	            }
	            storedFiles.push(f);
	            
	            var reader = new FileReader();
	            reader.onload = function (e) {
	                var html = "<div class='col-md-3'>"+
	                	"<img class='img-responsive' src='" + e.target.result + "' data-file='"+f.name+"'>" +
	                	"<p>Filename : "+f.name +"</p>"+
	                	"<a class='btn btn-sm btn-danger selFile' role='button'>Delete</a>"+
	                	 "</div>";
	                selDiv.append(html);
	                
	            }
	            reader.readAsDataURL(f); 
	        });
	        
	    }

	    function removeFile(e) {
	        var file = $(this).data("file");
	        for(var i=0;i<storedFiles.length;i++) {
	            if(storedFiles[i].name === file) {
	                storedFiles.splice(i,1);
	                break;
	            }
	        }
	        $(this).parent().remove();
	    }