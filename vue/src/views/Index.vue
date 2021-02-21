<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form path="misaakidis.cosmostriliza.cosmostriliza" type="game" :fields="[ ['creator', 1, 'string'] , ['rows', 2, 'string'] , ['cols', 3, 'string'] , ['strike', 4, 'string'] , ['reward', 5, 'string'] , ['gameState', 6, 'string'] , ['guest', 7, 'string'] , ['Oplayer', 8, 'string'] , ['Xplayer', 9, 'string'] , ['nextPlayer', 10, 'string'] ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
