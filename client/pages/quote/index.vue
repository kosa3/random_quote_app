<template>
  <div class="page">
    <div class="box quote">
      <p class="attention">下のカテゴリーを選択すると名言が一つ表示されるよ</p>
      <p>
        <nuxt-link to="/quote/life" class="tag">人生</nuxt-link>
        <nuxt-link to="/quote/like" class="tag">恋愛</nuxt-link>
        <nuxt-link to="/quote/die" class="tag">死</nuxt-link>
        <nuxt-link to="/quote/marriage" class="tag">結婚</nuxt-link>
        <nuxt-link to="/quote/happy" class="tag">幸福</nuxt-link>
        <nuxt-link to="/quote/money" class="tag">金</nuxt-link>
        <nuxt-link to="/quote/love" class="tag">愛</nuxt-link>
        <nuxt-link to="/quote/friend" class="tag">友達</nuxt-link>
        <nuxt-link to="/quote/woman" class="tag">女</nuxt-link>
        <nuxt-link to="/quote/time" class="tag">時間</nuxt-link>
        <nuxt-link to="/quote/broken_heart" class="tag">失恋</nuxt-link>
        <nuxt-link to="/quote/self" class="tag">自己</nuxt-link>
        <nuxt-link to="/quote/youth" class="tag">青春</nuxt-link>
        <nuxt-link to="/quote/man" class="tag">男</nuxt-link>
        <nuxt-link to="/quote/society" class="tag">処世</nuxt-link>
        <nuxt-link to="/quote/lonely" class="tag">孤独</nuxt-link>
        <nuxt-link to="/quote/destiny" class="tag">運命</nuxt-link>
        <nuxt-link to="/quote/truth" class="tag">真実</nuxt-link>
        <nuxt-link to="/quote/misfortune" class="tag">不幸</nuxt-link>
        <nuxt-link to="/quote/freedom" class="tag">自由</nuxt-link>
        <nuxt-link to="/quote/education" class="tag">教育</nuxt-link>
        <nuxt-link to="/quote/positive" class="tag">ポジティブ</nuxt-link>
        <nuxt-link to="/quote/management" class="tag">マネジメント</nuxt-link>
        <nuxt-link to="/quote/challenge" class="tag">挑戦</nuxt-link>
        <nuxt-link to="/quote/effort" class="tag">努力</nuxt-link>
        <nuxt-link to="/quote/future" class="tag">未来</nuxt-link>
      </p>
    </div>
  </div>
</template>
<style>
.box.quote {
  background-color: #ff7665;
  display: flex;
  flex-direction: column;
  align-items: center;
}
p {
  width: 100%;
  color: white;
  text-align: center;
  font-family: Menlo, sans-serif;
}
p.attention {
  font-weight: bold;
  font-size: 20px;
}
.tag {
  display: inline-block;
  font-size: 25px;
  font-weight: bold;
  margin-top: 20px;
  padding: 0.25em 0.5em;
  text-decoration: none;
  color: #fff;
  background: #00bcd4;
  transition: 0.4s;
}

.tag:hover {
  background: #1ec7bb;
}
</style>
