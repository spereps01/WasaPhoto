<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			ut: null,
            username: null,
			users: null,
            showModal: false,
			images : null,
			foto:null,
			idp: null,
			showComments:false,
			chI:false,
			response:null,
			likeStatus:false,
		}
	},
	methods: {
		load() {
			return load
		},

        async setMyUsername() {
	
			this.loading = true;
			this.errormsg = null;

			try {
				let u = await this.$axios.get("/search/" + localStorage.getItem("username"),{
                    	headers: {
							Authorization: localStorage.getItem("token")
						}
                });
				this.ut = u.data[0].Id
                

				let response = await this.$axios.put("/profile/" + this.ut,{
                    username: this.username
                });
				localStorage.setItem("username",this.username)
  
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			alert("Username changed correctly!")
		},




        
		async getUser() {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/search/" + localStorage.getItem("username"),{
                    	headers: {
							Authorization: localStorage.getItem("token")
						}
                });
				this.users = response.data;
   
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async getOneUser() {
			this.showModal = false;
			this.loading = true;
			this.errormsg = null;
			try {

				let response = await this.$axios.get("/search/" + localStorage.getItem("username"),{
                    	headers: {
							Authorization: localStorage.getItem("token")
						}
                });

				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.showComments=false;
		},

		async uploadPhoto() {
			const reader = new FileReader();
			console.log(this.$refs.file.files[0])
			reader.readAsArrayBuffer(this.$refs.file.files[0]);
			reader.onloadend = (event) => {
			this.images = new Uint8Array(event.target.result);
	


		}},


		async searchUser() {
			this.$router.push("/ricerca");
		},


		async UpPh() {
	
			this.loading = true;
			this.errormsg = null;
			try {
				
				let u = await this.$axios.get("/search/" + localStorage.getItem("username"),{
                    	headers: {
							Authorization: localStorage.getItem("token")
						}
                });
				this.ut = u.data[0].Id
				let response = await this.$axios.post("/profile/"+ this.ut +"/photo", this.images);
				this.foto = response.data;
				console.log(this.foto.Photo)
		
   
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser();
		},
		async deletePhoto(idp) {
	
			this.loading = true;
			this.errormsg = null;
			try {
			
				let response = await this.$axios.delete("/profile/"+localStorage.getItem("id").toString()+"/photo/"+idp.toString());
				this.foto = response.data;
		
   
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser();
		},
       async likePhoto(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.put("/photo/"+ id.toString()+"/like/"+localStorage.getItem("id").toString());
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
            this.getOneUser()
			if (this.likeStatus==false){
				this.likeStatus=true
			}
			else{
				alert("Like already added!")
			}
		},
		
		async unlikePhoto(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.delete("/photo/"+ id.toString()+"/like/"+localStorage.getItem("id").toString());
			
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
            this.getOneUser()
			if (this.likeStatus==true){
				this.likeStatus=false
			}
			else{
				alert("Like already removed!")
			}
		},
        async commentPhoto(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.put("/photo/"+id.toString()+"/comment/"+localStorage.getItem("id"),{
					comment: this.comment
				});
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser()
		},
		async getComments(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/photo/"+id.toString()+"/comments");
				this.comments = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.showComments = true;
		},
        async checkID(id){
			if( id.toString() == localStorage.getItem("id")){
				this.chI =true;
			}
			else{
				this.chI=false;
			}

		},
		async deleteComment(idp,idc) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.delete("/photo/"+idc.toString()+"/comment/"+idp.toString());
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser()
		},
		async chanU(){
			this.showModal = true;
			this.loading=true;

		}

	},

}
</script>


<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">MyProfile</h1>
			<div class="btn-group me-2">
				<input type="file" accept="image/*" class="btn btn-outline-primary" @change="uploadPhoto" ref="file">
				<button class="btn btn-success" @click="UpPh">Upload</button>
			</div>
			
			
		</div>

		<div class="mb-3">
			<a href="javascript:" class="btn btn-primary" @click="getOneUser()">Info</a>

            <label for="description" class="btn btn-warning" @click="chanU()">Change Username </label>

			<button class="btn btn-success" @click="searchUser()">Search</button>

            
		</div>
        <LoadingSpinner v-if="loading"></LoadingSpinner>
        



		<div class="card" v-if="!loading" v-for="u in users">
        

			<div class="card-body">
				<p class="card-text">
					Username:{{" "  + u.Username }}<br/>
					Numero di foto:{{" "  + u.N_p }}<br/>
					Numero di followers:{{" "  + u.N_followers }}<br/>
					Numero di followings:{{" "  + u.N_followings }}<br/>

				</p>
				<div v-if="!loading" v-for="p in u.Photos">
						<img :src="'data:image/png;base64,' + p.Photo" width=300 height=300 /><br/>
						<a href="javascript:" class="btn btn-danger" @click="deletePhoto(p.Id_photo)">Delete Photo</a>
						<div class="card-body">
							<a href="javascript:"  class="btn btn-primary" @click="likePhoto(p.Id_photo)">Like</a>
							<a href="javascript:" class="btn btn-danger" @click="unlikePhoto(p.Id_photo)">Unlike</a>
							<a href="javascript:" class="btn btn-warning" @click="getComments(p.Id_photo)">Show Comments</a>
							<input type="string" class="form-control" id="comment" v-model="comment" placeholder="enter the comment">
           					<a href="javascript:" class="btn btn-success" style="width: 160px; height: 35px;" @click="commentPhoto(p.Id_photo)">Send Comment</a>
						</div>
				</div>
				<a href="javascript:" class="btn btn-secondary" @click="loading = true">Close</a>
			</div>
		</div>



        <div v-if="showModal" class="modal-overlay">
            <div class="modal-content">
                <h2>Change Username</h2>
                <input type="string" class="form-control" id="username" v-model="username" placeholder="enter the new username">
                <div class="mb-3">
                <a href="javascript:" class="btn btn-primary" @click="setMyUsername()">Change</a>
                <a href="javascript:" class="btn btn-secondary" @click="showModal = false">Close</a>
                </div>
            </div>
        </div>

		<div v-if="showComments" class="modal-overlay">
            <div class="modal-content">
                <h2>Comments</h2>

                <div v-for="c in comments">
					user {{c.Owner_id}}
					comment: {{c.Comment}}
					<a  href="javascript:" class="btn btn-danger" v-if="checkID(c.Owner_id)==true">Delete Comment</a>
					<a  href="javascript:" class="btn btn-danger" v-if="chI==true" style="width: 170px; height: 35px;" @click="deleteComment(c.Comment_id,c.Id_photo)">Delete Comment</a>
					<p></p>
				</div>

				<a href="javascript:" class="btn btn-secondary" @click="getOneUser()">Close Comments</a>
            </div>
        </div>






	
    </div>
</template>