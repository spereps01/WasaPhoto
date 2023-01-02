<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			utente:null,
			users: [],
		}
	},
	methods: {
		load() {
			return load
		},

		async newItem() {
			this.$router.push("/ricerca/profile");
		},
		async getUser(utente) {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/search/" + utente);
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
			<h1 class="h2">Ricerca</h1>
			
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		
		<div class="mb-3">
			<label for="description" class="form-label">Username</label>
			<input type="string" class="form-control" id="utente" v-model="utente" placeholder="enter username">
			<a href="javascript:" class="btn btn-primary" @click="getUser(utente)">Ricerca</a>
		</div>
		<LoadingSpinner v-if="loading"></LoadingSpinner>



		<div class="card" v-if="!loading" v-for="u in users">

			<div class="card-body">
				<p class="card-text">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						{{ u.Username }}
					</button>

				</p>
			</div>
		</div>
		
	</div>
</template>

<style scoped>
.card {
	margin-bottom: 20px;
}
</style>
