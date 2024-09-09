<script lang="ts">
    import Top from "$lib/components/Top.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import OptionsList from "$lib/components/optionsList.svelte";
    import kickPlayer from '$lib/images/kickPlayer.svg';
    import operator from '$lib/images/operator.svg';
    import banPlayer from '$lib/images/banPlayer.svg';
    import banIP from '$lib/images/banIP.svg';

    interface Player{
        name: string;
        isOnline: boolean;
    }

    const players: Player[] = [{name: "player1", isOnline:false}, 
    {name: "player2", isOnline:true}, 
    {name: "player3", isOnline:false},
    {name: "player4", isOnline:true}, 
    {name: "player5", isOnline:false}]
</script>
<style>
    section{
        display: flex;
        flex-direction: column;
        height: 100vh;
        box-sizing: border-box;
        padding: 0;
        background-color: rgb(88, 88, 88);
        color: white;
    }

    section > div{
        width: 100%;
        flex: 1;
    }

    nav{
        display: inline-block;
        float: left;
        height: 100%;
        width: 240px;
    }

    div > section{
        width: 320px;
        height: 100%;
        float: right;
        box-sizing: border-box;
        padding-top: 35px;
    }

    a{
        box-sizing: border-box;
        height: 100px;
        margin-top: 60px;
        margin-right: 40px;
        padding-top: 30px;
        text-decoration: none;
        font-weight: bold;
        font-size: 30px;
        color: white;
        text-align: center;
    }

    .green{
        background-color: green;
    }

    .green:hover{
        background-color: rgb(26, 168, 26);
    }

    .red{
        background-color: rgb(218, 0, 0);
    }

    .red:hover{
        background-color: rgb(255, 0, 0);
    }

    main{
        display: inline-block;
        padding: 3%;
        width: calc(100% - 690px);
    }

    .playerWrapper{
        box-sizing: border-box;
        display: inline-block;
        width: 100%;
        height: 54px;
        margin-top: 18px;
        background-color: rgb(48, 48, 48);
    }

    .playerInfo{
        display: inline-block;
        background-color: rgb(48, 48, 48);
        color: white;
        font-weight: 600;
        font-size: 22px;
        padding: 13px;
    }

    .playerInfo::before{
        content: url($lib/images/steve.png);
        margin-right: 10px;
    }

    .buttonsContainer{
        display: inline-block;
        float: right;
        background-color: rgb(48, 48, 48);
        color: white;
        padding: 8px 6px 0px;
    }

    button{
        width: 38px;
        height: 38px;
        background-color: transparent;
        border: 0;
        color: white;
        margin: 0px 2px;
    }

    button.smaller{
        width: 36px; 
        height: 36px
    }
</style>

<section>
    <Top showServers={true} isLoggedIn={true}/>
    <div>
        <nav><OptionsList name={"name"} version={"version"} author={"author"}/></nav>
        <main>
            {#each players as player}
                <div class="playerWrapper">
                    <div class="playerInfo">{player.name}</div>
                    <div class="buttonsContainer">
                        <button><img src={kickPlayer} alt="kick"></button>
                        <button><img src={operator} alt="Operator"></button>
                        <button><img src={banPlayer} alt="Ban"></button>
                        <button class="smaller"><img src={banIP} alt="BanIP"></button>
                    </div>
                </div>
            {/each}
        </main>
        <section>
            <a href="/configuration/players/whitelist" class="green">Whitelist</a>
            <a href="#" class="green">Operators</a>
            <a href="#" class="red">Banned players</a>
            <a href="#" class="red">Banned IP's</a>
        </section>
    </div>
    <Footer/>
</section>