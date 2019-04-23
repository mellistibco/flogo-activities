package statistica

//
//#include <string.h>
//
//#define STRING_EQUAL(a,b) (strcasecmp(a,b)==0)
//#define STRING_NOT_EQUAL(a,b) (strcasecmp(a,b)!=0)
//#define STRING_SET(a,b)   strcpy(a,b)
//
//
//void _BTrees_Prediction_T14_19_26(
//      const char * _IsPersonal__,
//      const char * _IsWork__,
//      const char * _Contact__,
//      const char * _Domain__,
//      const char * _OrgName__,
//      char * _pRet
//   )
//   {
//     double _MaxValue;
//     int _i;
//     double _den;
//     double _LearningRate;
//     double _PredictProb[2];
//     _MaxValue = -1.0E30;
//     _den = 0;
//     _LearningRate = 0.100000;
//
//     for(_i=0;_i<2;_i++)
//     {
//         _PredictProb[_i] = 0;
//     }
//
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += 1.000000 * -1.00000000000000e+000;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += 1.000000 * 1.00000000000000e+000;
//
//    }
//    else {
//    _PredictProb[0]  += 1.000000 * 2.24489795918367e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.67667641618306e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.67667641618306e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.28910686284502e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.60405162316125e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.60405162316125e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.29324268226798e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.54000512520268e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.54000512520268e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.46979727811500e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.48336812558526e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.48336812558526e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.01222057969652e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.43316172868151e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.43316172868151e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.28420186314290e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.38856012939146e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.38856012939146e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.43013496031380e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.34886210463064e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.34886210463064e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.41908586449384e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.31346869115432e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.31346869115432e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.70790065072818e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.28186552980162e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.28186552980162e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.43190273178864e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.25360876948171e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.25360876948171e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.91040318890244e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.22831371508128e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.22831371508128e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.01088989041592e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.20564561097666e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.20564561097666e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.47043040731773e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.18531210173376e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.18531210173376e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.48151774335250e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.16705702112410e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.16705702112410e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.47120118401587e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.15065524150540e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.15065524150540e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.71688508050180e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.13590837599054e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.13590837599054e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.23261801023773e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.12264117129841e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.12264117129841e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.41314239208232e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.11069846372513e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.11069846372513e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.70356615457504e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.09994259713515e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.09994259713515e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.44433814284289e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.09025122229709e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.09025122229709e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.50822258438432e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.08151541277196e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.08151541277196e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.11729808865499e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.07363804499538e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.07363804499538e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.31538764129511e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.06653239999490e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.06653239999490e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.68884413333164e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.06012095195323e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.06012095195323e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.49647548437505e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.05433431502603e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.05433431502603e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.28835972735957e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.04911032479435e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.04911032479435e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.51473309743832e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.04439323474155e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.04439323474155e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.07102917558630e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.04013301139471e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.04013301139471e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.40014251220880e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.03628471441966e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.03628471441966e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.37353219484173e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.03280795012669e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.03280795012669e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.31290642177219e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.02966638862796e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.02966638862797e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.35414095078447e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.02682733636180e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.02682733636180e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.00536546727232e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.02426135692190e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.02426135692191e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.30418769800387e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.02194193414986e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.02194193414986e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.28270087915915e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.01984517230244e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.01984517230244e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.07568110835057e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.01794952882277e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.01794952882277e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.56177229892950e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.01623557584992e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.01623557584992e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.91787801553943e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.01468578711386e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.01468578711386e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.54999378874440e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.01328434729893e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.01328434729893e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.08129662392719e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.01201698133101e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.01201698133102e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.11275814787576e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.01087080136181e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.01087080136181e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.38598554080224e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00983416949769e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00983416949769e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.00983416949799e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00889657455629e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00889657455629e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.25222414363903e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00804852133806e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00804852133806e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.23148734131275e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00728143107713e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00728143107713e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.34546905179407e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00658755188948e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00658755188948e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.59784709102849e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00595987817099e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00595987817099e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.59764676962900e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00539207801518e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00539207801518e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.86475391141754e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00487842782346e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00487842782347e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.27575332473952e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00441375337130e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00441375337131e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.07237437572255e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00399337667318e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00399337667318e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.00638940267840e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00361306805985e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00361306805985e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.16363094605980e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00326900294333e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00326900294333e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.55273865608565e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00295772280038e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00295772280038e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.27193840410166e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00267609995404e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00267609995404e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.53143145916979e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00242130577623e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00242130577623e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.07895361497128e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00219078197321e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00219078197321e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.69567812611116e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00198221465021e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00198221465021e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.11155160325559e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00179351088221e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00179351088221e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.69506693981867e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00162277754545e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00162277754545e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.59142542921891e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00146830218894e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00146830218894e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.86329211257994e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00132853574713e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00132853574713e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.85234390212838e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00120207691473e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00120207691473e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.53098022762677e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00108765802252e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00108765802252e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.03796158940006e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00098413226858e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00098413226858e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.00157461162200e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00089046217409e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00089046217409e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.15405164511706e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00080570914526e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00080570914526e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.32674437181360e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00072902403497e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00072902403496e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.03463359117823e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00065963860757e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00065963860757e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.82469347991421e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00059685782025e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00059685782025e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.33432809636718e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00054005284227e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00054005284227e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.27920792049333e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00048865474139e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00048865474139e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.66682955158109e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00044214877333e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00044214877333e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.94143655809248e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00040006921663e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00040006921663e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.52186089063133e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00036199470054e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00036199470054e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.74431232373241e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00032754397879e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00032754397878e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.00006550879323e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00029637210655e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00029637210655e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.53855272988380e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00026816698212e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00026816698212e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.20006436008039e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00024264621831e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00024264621831e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.00038823391026e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00021955431219e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00021955431219e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.25005488857518e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00019866008462e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00019866008462e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.17025926087261e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00017975436383e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00017975436383e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.66672658478407e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00016264788972e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00016264788972e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.64943668887417e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00014716941785e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00014716941785e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.50004415083255e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00013316400410e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00013316400410e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.25003329100337e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00012049145275e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00012049145275e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.53849861275406e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00010902491226e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00010902491226e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.22644183629840e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00009864960501e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00009864960501e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.53849189218009e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00008926167777e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00008926167777e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.80004998654787e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00008076716186e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00008076716186e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.00001615343186e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00007308103215e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00007308103215e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.66669102699545e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00006612635578e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00006612635578e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.41511305462858e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00005983352190e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00005983352190e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.60379277678475e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00005413954468e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00005413954468e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.86276526766106e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00004898743278e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00004898743278e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.08696717116138e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00004432561876e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00004432561877e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.00000886511722e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00004010744288e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00004010744288e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.09375877351402e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00003629068595e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00003629068595e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.45834391811371e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00003283714674e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00003283714674e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.45834291085744e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00002971225956e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00002971225956e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.00004753991120e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00002688474825e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00002688474825e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.29630326639444e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00002432631311e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00002432631311e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.86277848700986e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00002201134763e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00002201134763e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.45833975334259e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00001991668219e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00001991668219e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.36364179542876e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00001802135211e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00001802135211e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.00000360427858e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00001630638784e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00001630638784e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.17392013256439e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00001475462506e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00001475462506e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.38298280452428e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00001335053290e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00001335053290e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.66667111675148e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00001208005849e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00001208005849e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.25000301998774e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00001093048630e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00001093048630e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.52175338782279e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000989031084e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000989031083e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.42857425428706e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000894912155e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000894912155e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.82354520371677e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000809749859e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000809749859e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.86275621317579e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000732691853e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000732691853e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.53846379291637e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000662966907e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000662966907e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.61017162651232e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000599877185e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000599877185e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.66666866642464e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000542791258e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000542791258e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.18181946483592e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000491137787e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000491137787e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.25000614018925e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000444399804e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000444399804e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.55556049560945e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000402109535e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000402109535e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.44681450056857e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000363843724e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000363843724e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.60000116442642e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000329219392e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000329219392e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.57143092713767e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000297890005e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000297890005e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.72727375657709e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000269542007e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000269542007e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.69231183632269e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000243891681e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000243891681e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.52173987276366e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000220682308e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000220682308e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.27451036611084e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000199681601e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000199681601e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.00000159343597e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000180679377e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000180679377e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.69565531681911e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000163485455e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000163485455e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.00000065404078e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000147927752e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000147927752e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.82353202577702e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000133850561e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000133850561e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.42857181028128e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000121112993e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000121112993e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.10169518289302e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000109587565e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000109587565e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.60377393631799e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000099158927e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000099158927e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.86274646315283e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000089722706e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000089722706e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.54545535568097e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000081184460e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000081184460e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.85185216900661e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000073458736e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000073458736e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.27450999051348e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000066468212e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000066468212e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.31914962597255e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000060142925e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000060142925e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.44680940367316e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000054419568e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000054419568e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.34615399007381e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000049240861e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000049240861e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.81818248368899e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000044554973e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000044554973e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.71698155206033e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000040315006e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000040315006e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.15384624551892e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000036478526e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000036478526e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.36363646191217e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000033007135e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000033007135e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.29629638230959e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000029866091e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000029866091e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.56250009071759e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000027120731e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000027064359e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.05469241346412e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000024452287e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000024452287e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.44680888428788e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000022125344e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000022163750e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.59945431906256e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000020113254e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000020019839e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.82086430632994e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000018161080e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000018114699e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.17606088472971e-001;
//
//    }
//    _PredictProb[0]  += _LearningRate * 1.44376839142293e-001;
//
//    _PredictProb[0]  += _LearningRate * 1.66862320568650e-001;
//
//    _PredictProb[0]  += _LearningRate * 6.39493433394050e-002;
//
//    _PredictProb[0]  += _LearningRate * 1.06025534940238e-001;
//
//    _PredictProb[0]  += _LearningRate * 1.53835356610398e-002;
//
//    _PredictProb[0]  += _LearningRate * 1.04713668882529e-001;
//
//    _PredictProb[0]  += _LearningRate * -1.04972059599434e-002;
//
//    _PredictProb[0]  += _LearningRate * 1.39007453829224e-002;
//
//    _PredictProb[0]  += _LearningRate * -1.18512217126737e-002;
//
//    _PredictProb[0]  += _LearningRate * 1.10980524003574e-001;
//
//    _PredictProb[0]  += _LearningRate * 1.30892591883103e-001;
//
//    _PredictProb[0]  += _LearningRate * 1.03981195204582e-001;
//
//    _PredictProb[0]  += _LearningRate * 6.49449751592557e-002;
//
//    _PredictProb[0]  += _LearningRate * -2.86915697371743e-002;
//
//    _PredictProb[0]  += _LearningRate * 6.82897458253107e-002;
//
//    _PredictProb[0]  += _LearningRate * 6.63902221863753e-003;
//
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[0]  += _LearningRate * -5.00000018817760e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[0]  += _LearningRate * 5.00000012954867e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -5.64077844897715e-002;
//
//    }
//    _PredictProb[0]  += _LearningRate * 7.61747903917490e-002;
//
//    _PredictProb[0]  += _LearningRate * 1.18161667128037e-002;
//
//    _PredictProb[0]  += _LearningRate * -2.58714554222625e-002;
//
//
//   if(_MaxValue <= _PredictProb[0])
//   {
//     _MaxValue = _PredictProb[0];
//     STRING_SET(_pRet,"No");
//   }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += 1.000000 * -1.00000000000000e+000;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += 1.000000 * 1.00000000000000e+000;
//
//    }
//    else {
//    _PredictProb[1]  += 1.000000 * 5.45454545454545e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.67667641618306e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.67667641618306e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.05836339962735e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.60405162316125e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.60405162316125e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.71552600709018e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.54000512520268e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.54000512520268e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.56792597883095e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.48336812558526e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.48336812558526e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.75467780018729e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.43316172868151e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.43316172868151e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.13717338507288e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.38856012939146e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.38856012939146e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.99576301088572e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.34886210463064e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.34886210463064e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.72543938859053e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.31346869115432e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.31346869115432e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.44912782486027e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.28186552980162e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.28186552980162e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -8.96920561664423e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.25360876948171e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.25360876948171e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.27107429908535e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.22831371508128e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.22831371508128e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.04566274301625e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.20564561097666e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.20564561097666e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.80604031401232e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.18531210173376e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.18531210173376e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.35056280640522e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.16705702112410e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.16705702112410e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.49051289594226e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.15065524150540e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.15065524150540e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.33535506261251e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.13590837599054e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.13590837599054e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.42057040186973e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.12264117129841e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.12264117129841e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.33634117512132e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.11069846372513e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.11069846372513e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.22656763129403e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.09994259713515e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.09994259713515e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.99992121175411e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.09025122229709e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.09025122229709e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.51814510138686e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.08151541277196e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.08151541277196e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.53627210153570e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.07363804499538e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.07363804499538e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.05700792604069e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.06653239999490e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.06653239999490e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.55949509432875e-003;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.06012095195323e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.06012095195323e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -8.43353491992218e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.05433431502603e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.05433431502603e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -5.61592701669562e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.04911032479435e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.04911032479435e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.68303677493144e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.04439323474155e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.04439323474155e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.37574360947498e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.04013301139471e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.04013301139471e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.23625904053076e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.03628471441966e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.03628471441966e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.75121199473552e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.03280795012669e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.03280795012669e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -5.35405101077297e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.02966638862796e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.02966638862797e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.11770364191731e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.02682733636180e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.02682733636180e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.37115420553246e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.02426135692190e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.02426135692191e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.62092600261607e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.02194193414986e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.02194193414986e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.18495161179167e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.01984517230244e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.01984517230244e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 6.27480646537856e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.01794952882277e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.01794952882277e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.07527489903348e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.01623557584992e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.01623557584992e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.62521679829206e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.01468578711386e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.01468578711386e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.68624771615059e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.01328434729893e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.01328434729893e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.27789601009584e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.01201698133101e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.01201698133102e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.17302525094970e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.01087080136181e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.01087080136181e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.15635480031423e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00983416949769e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00983416949769e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.16896130621610e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00889657455629e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00889657455629e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.45074825388461e-003;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00804852133806e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00804852133806e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.36583141491027e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00728143107713e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00728143107713e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.11272920690614e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00658755188948e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00658755188948e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.00131751037796e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00595987817099e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00595987817099e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.94108648337246e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00539207801518e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00539207801518e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.11230935066991e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00487842782346e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00487842782347e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.66829280927429e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00441375337130e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00441375337131e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.85348657532144e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00399337667318e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00399337667318e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.44149693710577e-003;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00361306805985e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00361306805985e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.32748918132206e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00326900294333e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00326900294333e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.02107530672549e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00295772280038e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00295772280038e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.66765257426692e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00267609995404e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00267609995404e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.03829126602800e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00242130577623e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00242130577623e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.00193704462164e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00219078197321e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00219078197321e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.12294078779005e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00198221465021e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00198221465021e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.04692651004308e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00179351088221e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00179351088221e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.61883367477360e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00162277754545e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00162277754545e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -5.00162277754653e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00146830218894e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00146830218894e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.32692016180550e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00132853574713e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00132853574713e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -8.00212565719751e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00120207691473e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00120207691473e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.09934596468633e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00108765802252e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00108765802252e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.36393299764332e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00098413226858e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00098413226858e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.04187169422338e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00089046217409e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00089046217409e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.17042117199823e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00080570914526e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00080570914526e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.94165041714906e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00072902403497e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00072902403496e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.27137178577299e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00065963860757e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00065963860757e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.08709992143461e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00059685782025e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00059685782025e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.29047660846973e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00054005284227e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00054005284227e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.34629924499719e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00048865474139e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00048865474139e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.15395892032426e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00044214877333e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00044214877333e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.14348878397120e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00040006921663e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00040006921663e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.11120001538128e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00036199470054e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00036199470054e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.07150614172191e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00032754397879e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00032754397878e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.54253942916802e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00029637210655e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00029637210655e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -8.82405242131245e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00026816698212e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00026816698212e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.16689013910247e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00024264621831e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00024264621831e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.80859840395479e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00021955431219e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00021955431219e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.66673985144011e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00019866008462e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00019866008462e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.63168350528928e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00017975436383e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00017975436383e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.45838576168981e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00016264788972e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00016264788972e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.07846645346996e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00014716941785e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00014716941785e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.59095591753494e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00013316400410e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00013316400410e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.16677763664525e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00012049145275e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00012049145275e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.25003012286384e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00010902491226e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00010902491226e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.04653444707552e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00009864960501e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00009864960501e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.84622973053201e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00008926167777e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00008926167777e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.66669642054850e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00008076716186e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00008076716186e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.57462274563420e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00007308103215e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00007308103215e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.12246538555414e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00006612635578e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00006612635578e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.44690699670458e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00005983352190e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00005983352190e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.17393905811194e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00005413954468e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00005413954468e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.44446008477061e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00004898743278e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00004898743278e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.53847661153646e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00004432561876e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00004432561877e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -8.49064130772064e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00004010744288e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00004010744288e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.25001002685000e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00003629068595e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00003629068595e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.36364626109923e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00003283714674e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00003283714674e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.40745605478660e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00002971225956e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00002971225956e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.09096311290554e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00002688474825e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00002688474825e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.39345011544317e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00002432631311e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00002432631311e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -5.31917481531395e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00002201134763e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00002201134763e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.57451023414952e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00001991668219e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00001991668219e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.09094530630057e-003;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00001802135211e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00001802135211e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.80000648767937e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00001630638784e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00001630638784e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.25000407665780e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00001475462506e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00001475462506e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.11111438986159e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00001335053290e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00001335053290e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.22641836909677e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00001208005849e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00001208005849e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.36365173853839e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00001093048630e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00001093048630e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.34615678895519e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000989031084e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000989031083e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.16981561285371e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000894912155e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000894912155e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.66666964976364e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000809749859e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000809749859e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.44828144561469e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000732691853e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000732691853e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.57143380473999e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000662966907e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000662966907e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.32653237112764e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000599877185e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000599877185e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.66666866608993e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000542791258e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000542791258e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.25000135692195e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000491137787e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000491137787e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.25000122775238e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000444399804e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000444399804e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.66666814779430e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000402109535e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000402109535e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.08695739578654e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000363843724e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000363843724e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.34782924941629e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000329219392e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000329219392e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.14035491613769e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000297890005e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000297890005e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.00000059552950e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000269542007e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000269542007e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.00000323731564e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000243891681e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000243891681e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.57447275484339e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000220682308e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000220682308e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.88888972288407e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000199681601e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000199681601e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.20689703409492e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000180679377e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000180679377e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.60377597057235e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000163485455e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000163485455e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.61538775820110e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000147927752e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000147927752e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.76190616713769e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000133850561e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000133850561e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.86274559737366e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000121112993e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000121112993e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.18181846731161e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000109587565e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000109587565e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.42857174177597e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000099158927e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000099158927e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.18367528662782e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000089722706e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000089722706e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -8.92857303606770e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000081184460e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000081184460e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -8.33333467037165e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000073458736e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000073458736e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.27450999182518e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000066468212e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000066468212e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.00000013461685e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000060142925e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000060142925e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.45161307808816e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000054419568e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000054419568e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.94915275688890e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000049240861e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000049240861e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.14285785224613e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000044554973e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000044554973e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.00000008914337e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000040315006e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000040315006e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -8.92857216240831e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000036478526e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000036478526e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.41509444406735e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000033007135e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000033007135e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.44827608501791e-002;
//
//    }
//    if ( _Domain__ != NULL &&  STRING_EQUAL(_Domain__,"org")  || _Domain__ != NULL &&  STRING_EQUAL(_Domain__,"edu")  || _Domain__ != NULL &&  STRING_EQUAL(_Domain__,"gov")  || _Domain__ != NULL &&  STRING_EQUAL(_Domain__,"net")  || _Domain__ != NULL &&  STRING_EQUAL(_Domain__,"uk")  || _Domain__ != NULL &&  STRING_EQUAL(_Domain__,"com")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000029866091e-001;
//
//    }
//    else if ( _Domain__ != NULL &&  STRING_EQUAL(_Domain__,"jp")  || _Domain__ != NULL &&  STRING_EQUAL(_Domain__,"foo")  || _Domain__ != NULL &&  STRING_EQUAL(_Domain__,"boo")  || _Domain__ != NULL &&  STRING_EQUAL(_Domain__,"goo")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000029866091e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.66666676879527e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000027023956e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000027023956e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.45833341058363e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000024501008e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000024606413e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.12659911380095e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000022159367e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000022125344e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.36702309271886e-001;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000020049770e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000020064704e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.98293916196745e-002;
//
//    }
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000018143373e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000018153674e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.61388601361293e-002;
//
//    }
//    _PredictProb[1]  += _LearningRate * -1.73913048529614e-001;
//
//    if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"Yes")  ) {
//        _PredictProb[1]  += _LearningRate * -5.00000015877368e-001;
//
//    }
//    else if ( _IsWork__ != NULL &&  STRING_EQUAL(_IsWork__,"No")  ) {
//        _PredictProb[1]  += _LearningRate * 5.00000016920954e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.54363811410209e-002;
//
//    }
//    _PredictProb[1]  += _LearningRate * -8.29959681870527e-002;
//
//    _PredictProb[1]  += _LearningRate * 2.19188237220562e-002;
//
//    _PredictProb[1]  += _LearningRate * -1.16489350050873e-001;
//
//    _PredictProb[1]  += _LearningRate * -7.90538376484225e-002;
//
//    _PredictProb[1]  += _LearningRate * -7.90260143656756e-002;
//
//    _PredictProb[1]  += _LearningRate * -6.60607809096840e-004;
//
//    _PredictProb[1]  += _LearningRate * 4.66850380953601e-002;
//
//    _PredictProb[1]  += _LearningRate * -7.34970946342996e-002;
//
//    _PredictProb[1]  += _LearningRate * -1.37418741833882e-001;
//
//    _PredictProb[1]  += _LearningRate * -6.26154485775515e-002;
//
//    _PredictProb[1]  += _LearningRate * -1.62928011459039e-002;
//
//    _PredictProb[1]  += _LearningRate * -4.48417901565694e-002;
//
//    _PredictProb[1]  += _LearningRate * -4.91988323646397e-002;
//
//    _PredictProb[1]  += _LearningRate * 4.47778116133455e-002;
//
//    _PredictProb[1]  += _LearningRate * -5.38269056998415e-002;
//
//    _PredictProb[1]  += _LearningRate * 6.55514165785107e-002;
//
//    _PredictProb[1]  += _LearningRate * -7.68249124601082e-002;
//
//    _PredictProb[1]  += _LearningRate * 1.00126862505496e-001;
//
//
//   if(_MaxValue <= _PredictProb[1])
//   {
//     _MaxValue = _PredictProb[1];
//     STRING_SET(_pRet,"Yes");
//   }
//
//}
//
import "C"

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log Is the Default package logger
var log = logger.GetLogger("activity-tibco-inference")

// InferfenceActivity is an Activity that is used to Invoke a ML Model using flogo-ml framework
type ModelActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a New InferfenceActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ModelActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *ModelActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Runs an ML model
func (a *ModelActivity) Eval(context activity.Context) (done bool, err error) {

	var i1 = C.CString(context.GetInput("_IsPersonal__").(string))
	var i2 = C.CString(context.GetInput("_IsWork__").(string))
	var i3 = C.CString(context.GetInput("_Contact__").(string))
	var i4 = C.CString(context.GetInput("_Domain__").(string))
	var i5 = C.CString(context.GetInput("_OrgName__").(string))

	var result C.char
	C._BTrees_Prediction_T14_19_26(i1, i2, i3, i4, i5, &result)

	context.SetOutput("result", C.GoString(&result))

	return true, nil
}
