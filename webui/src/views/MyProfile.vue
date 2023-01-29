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
				let u = await this.$axios.get("/search/" + this.$route.params.username);
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
				let response = await this.$axios.get("/search/" + this.$route.params.username);
				this.users = response.data;
   
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
			
		</div>

		<div class="mb-3">
			<a href="javascript:" class="btn btn-primary" @click="getUser()">Info</a>

            <label for="description" class="btn btn-warning" @click="showModal = true">Change Username </label>

			<label for="description" class="btn btn-secondary" >Upload Photo </label>
            
		</div>
        <LoadingSpinner v-if="loading"></LoadingSpinner>
        



		<div class="card" v-if="!loading" v-for="u in users">
        

			<div class="card-body">
				<p class="card-text">
					Username:{{" "  + u.Username }}<br/>
					Numero di foto:{{" "  + u.N_p }}<br/>
					Foto:{{" "  + u.Photos }}<br/>
					Numero di followers:{{" "  + u.N_followers }}<br/>
					Numero di followings:{{" "  + u.N_followings }}<br/>


				</p>
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