import { Args, Mutation, Query, Resolver } from "@nestjs/graphql";
import { GraphQLService } from "./graphql.service";
import { FootballAPIPlayer } from "../football/footballApiPlayer";
import { FootballService } from "src/football/football.service";

@Resolver()
export class GraphQLResolver {

  constructor(
    private readonly graphqlService: GraphQLService,
    private readonly footballService: FootballService
  ) {}

  @Query(() => String)
  status(): string {
    return 'OK';
  }

  @Query(() => [FootballAPIPlayer], { name: "footballAPIPlayers"})
  footballAPIPlayers(
    @Args("season") season: number,
    @Args("teamId") teamId: number,
    @Args("page") page: number
  ): Promise<FootballAPIPlayer[]> {
    return this.footballService.fetchPlayers(season, teamId, page);
  }

  @Mutation(() => String)
  async createSquad(
    @Args('teamId') teamId: string,
    @Args('playerId') playerId: string
  ): Promise<string> {
    const result = await this.graphqlService.createSquad(teamId, playerId);
    return result.squadId;
  }

  @Mutation(() => String)
  async createPlayer(
    @Args('footballapiPlayerId') footballapiPlayerId: string,
    @Args('footballapiTeamId') footballapiTeamId: string,
    @Args('playerName') playerName: string,
    @Args('teamAuth0UserId') teamAuth0UserId: string,
    @Args('birthDate') birthDate: string,
    @Args('nationality') nationality: string,
    @Args('height') height: string,
    @Args('weight') weight: string
  ): Promise<string> {
    const result = await this.graphqlService.createPlayer(
      footballapiPlayerId,
      footballapiTeamId,
      playerName,
      teamAuth0UserId,
      birthDate,
      nationality,
      height,
      weight
    )
    return result.playerId;
  }
}