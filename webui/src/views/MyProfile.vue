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
				let u = await this.$axios.get("/search/" + localStorage.getItem("username"));
				this.ut = u.data[0].Id
                

				let response = await this.$axios.put("/profile/" + this.ut,{
                    username: this.username
                });
  
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},




        
		async getUser() {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/search/" + localStorage.getItem("username"));
				this.users = response.data;
   
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async getOneUser() {
			this.loading = true;
			this.errormsg = null;
			try {

				let response = await this.$axios.get("/search/" + localStorage.getItem("username"));

				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},

		async uploadPhoto() {
			const reader = new FileReader();
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
				
				let u = await this.$axios.get("/search/" + localStorage.getItem("username"));
				this.ut = u.data[0].Id
				let response = await this.$axios.post("/profile/"+ this.ut +"/photo", this.images);
				this.foto = response.data;
				console.log(this.foto.Photo)
		
   
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},

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

            <label for="description" class="btn btn-warning" @click="showModal = true">Change Username </label>

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
						<a href="javascript:" class="btn btn-primary">Like</a>
						<a href="javascript:" class="btn btn-secondary">Comment</a>
						<a href="javascript:" class="btn btn-danger">Delete</a>
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






	
    </div>
</template>