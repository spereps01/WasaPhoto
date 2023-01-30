<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			ut: null,
			users: null,
			id:null,
			likeStatus: false,
			followStatus: false,
			banStatus: false,
		}
	},
	methods: {
		load() {
			return load
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
		async getOneUser() {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/search/" + this.$route.params.username);

				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async likePhoto(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.put("/photo/"+ id.toString()+"/like/"+localStorage.getItem("id").toString());
				this.likeStatus = true;
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser()
		},
		async unlikePhoto(id) {
	
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.delete("/photo/"+ id.toString()+"/like/"+localStorage.getItem("id").toString());
				this.likeStatus = false;
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser()
		},
		async followUser(id) {
	
			this.loading = true;
			this.errormsg = null;
			console.log(id)
			try {
				let response = await this.$axios.put("/users/"+ localStorage.getItem("id").toString()+"/follow/"+id.toString());
				this.followStatus = true;
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser()
		},
		async unfollowUser(id) {
	
			this.loading = true;
			this.errormsg = null;
			console.log(id)
			try {
				let response = await this.$axios.delete("/users/"+ localStorage.getItem("id").toString()+"/follow/"+id.toString());
				this.followStatus = false;
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser()
		},
		async banUser(id) {
	
			this.loading = true;
			this.errormsg = null;
			console.log(id)
			try {
				let response = await this.$axios.put("/users/"+ localStorage.getItem("id").toString()+"/ban/"+id.toString());
				this.banStatus = true;
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser()
		},
		async unbanUser(id) {
	
			this.loading = true;
			this.errormsg = null;
			console.log(id)
			try {
				let response = await this.$axios.delete("/users/"+ localStorage.getItem("id").toString()+"/ban/"+id.toString());
				this.banStatus = false;
				this.users = [response.data[0]];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.getOneUser()
		},
		

	},

	

}
</script>


<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Profilo</h1>
			
		</div>
		<div class="mb-3">
			<a href="javascript:" class="btn btn-info" @click="getOneUser()">Informations</a>
		</div>
		<LoadingSpinner v-if="loading"></LoadingSpinner>



		<div class="card" v-if="!loading" v-for="u in users">
			<a href="javascript:" class="btn btn-warning" v-if="followStatus == false" @click="followUser(u.Id)">Follow</a>
			<a href="javascript:" class="btn btn-warning" v-if="followStatus == true" @click="unfollowUser(u.Id)">Unfollow</a>
			<a href="javascript:" class="btn btn-danger" v-if="banStatus == false" @click="banUser(u.Id)">Ban</a>
			<a href="javascript:" class="btn btn-danger" v-if="banStatus == true" @click="unbanUser(u.Id)">Unban</a>

			<div class="card-body">
				<p class="card-text">
					Username:{{" "  + u.Username }}<br/>
					Numero di foto:{{" "  + u.N_p }}<br/>
					Numero di followers:{{" "  + u.N_followers }}<br/>
					Numero di followings:{{" "  + u.N_followings }}<br/>
				</p>
				<div v-if="!loading" v-for="p in u.Photos">
						<img :src="'data:image/png;base64,' + p.Photo" width=300 height=300 /><br/>
						<a href="javascript:" class="btn btn-primary" v-if="likeStatus == false" @click="likePhoto(p.Id_photo)">Like</a>
						<a href="javascript:" class="btn btn-primary" v-if="likeStatus == true" @click="unlikePhoto(p.Id_photo)">Unlike</a>
						<a href="javascript:" class="btn btn-secondary">Comment</a>
						

				</div>
			</div>
		</div>
    </div>
</template>